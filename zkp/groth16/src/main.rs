use bellman::{
    groth16::{
        create_random_proof, generate_random_parameters, prepare_verifying_key, verify_proof,
    },
    Circuit, ConstraintSystem, SynthesisError,
};
use bls12_381::{Bls12, Scalar};
use ff::Field;
use rand::thread_rng;

// Define our circuit struct
#[derive(Clone)]
struct CubicEquationCircuit {
    x: Option<Scalar>,
}

// Implement the Circuit trait for our cubic equation
impl Circuit<Scalar> for CubicEquationCircuit {
    fn synthesize<CS: ConstraintSystem<Scalar>>(self, cs: &mut CS) -> Result<(), SynthesisError> {
        // Allocate the `x` variable
        let x = cs.alloc(|| "x", || self.x.ok_or(SynthesisError::AssignmentMissing))?;

        // x^2
        let x_squared = cs.alloc(
            || "x^2",
            || {
                let x_val = self.x.ok_or(SynthesisError::AssignmentMissing)?;
                Ok(x_val.square())
            },
        )?;

        // Enforce x * x = x^2
        cs.enforce(
            || "x * x = x^2",
            |lc| lc + x,
            |lc| lc + x,
            |lc| lc + x_squared,
        );

        // x^3
        let x_cubed = cs.alloc(
            || "x^3",
            || {
                let x_val = self.x.ok_or(SynthesisError::AssignmentMissing)?;
                Ok(x_val.cube())
            },
        )?;

        // Enforce x * x^2 = x^3
        cs.enforce(
            || "x * x^2 = x^3",
            |lc| lc + x,
            |lc| lc + x_squared,
            |lc| lc + x_cubed,
        );

        // Enforce x^3 + x + 5 = 35
        cs.enforce(
            || "x^3 + x + 5 = 35",
            |lc| lc + x_cubed + x + (Scalar::from(5), CS::one()),
            |lc| lc + CS::one(),
            |lc| lc + (Scalar::from(35), CS::one()),
        );

        Ok(())
    }
}

fn main() {
    // Create an instance of our circuit (with the witness)
    let circuit = CubicEquationCircuit {
        x: Some(Scalar::from(3)), // The solution is x = 3
    };

    // Generate the parameters for our circuit
    let params =
        generate_random_parameters::<Bls12, _, _>(circuit.clone(), &mut thread_rng()).unwrap();

    // Create a proof with our parameters
    let proof = create_random_proof(circuit, &params, &mut thread_rng()).unwrap();

    let pvk = prepare_verifying_key::<Bls12>(&params.vk);
    // Verify the proof
    assert!(verify_proof(&pvk, &proof, &[]).is_ok());

    println!("Proof verified successfully!");
}
