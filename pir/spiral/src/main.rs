use bzip2_rs::DecoderReader;
use spiral_rs::client::{Client, PublicParameters};
use spiral_rs::key_value::{extract_result_impl, row_from_key};
use spiral_rs::params::Params;
use spiral_rs::util::params_from_json;
use spiral_server::db::sparse_db::SparseDb;
use spiral_server::db::write::update_database;
use spiral_server::server::process_query;
use std::io::Error;
use std::io::Read;

struct ServerState {
    params: &'static Params,
    db: SparseDb,
    rows: Vec<Vec<u8>>,
}

fn get_params() -> Params {
    let cfg_expand = r#"{
        "n": 2,
        "nu_1": 9,
        "nu_2": 5,
        "p": 256,
        "q2_bits": 22,
        "t_gsw": 7,
        "t_conv": 3,
        "t_exp_left": 5,
        "t_exp_right": 5,
        "instances": 4,
        "db_item_size": 32768
    }"#;
    params_from_json(cfg_expand)
}

fn main() -> Result<(), Error> {
    let params = get_params();

    // Server: init server state
    let db = SparseDb::new();
    let rows = vec![Vec::new(); params.num_items()];
    let mut server_state = ServerState {
        params: Box::leak(Box::new(params.clone())),
        db,
        rows,
    };

    // Server: update server state
    let kv_pairs: Vec<(&str, &[u8])> = vec![
        ("key-1", "value-1".as_bytes()),
        ("key-2", "value-2".as_bytes()),
        ("key-3", "value-3".as_bytes()),
    ];
    update_database(
        server_state.params,
        &kv_pairs,
        &mut server_state.rows,
        &mut server_state.db,
    );

    // Client: get Params from server, init client, generate keys and query
    let mut client = Client::init(&params);
    let public_params = client.generate_keys();
    let pp_serialized = public_params.serialize();
    let idx = row_from_key(&params, "key-1");
    let query = client.generate_query(idx);

    // Server: process query
    let pp = PublicParameters::deserialize(&params, &pp_serialized);
    let response = process_query(&params, &pp, &query, &server_state.db);

    // Client: decode, decompress and extract result
    let decrypted = client.decode_response(response.as_slice());
    let decompressed = decompress(&decrypted)?; // equal to server_state.rows[idx]
    let result = extract_result_impl("key-1", decompressed.as_slice()).unwrap();
    println!("result: {:?}", String::from_utf8(result).unwrap());
    Ok(())
}

/// Decompress the given data using bzip2.
fn decompress(data: &[u8]) -> Result<Vec<u8>, Error> {
    let mut decoder = DecoderReader::new(data);
    let mut decompressed = Vec::new();
    decoder.read_to_end(&mut decompressed)?;
    Ok(decompressed)
}
