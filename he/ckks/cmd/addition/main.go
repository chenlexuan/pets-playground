package main

import (
	"fmt"
	"math/rand"

	"github.com/tuneinsight/lattigo/v5/core/rlwe"
	"github.com/tuneinsight/lattigo/v5/he/hefloat"
)

func main() {
	// Instantiate the hefloat.Parameters using the HEFloatComplexParamsN12QP109 parameters literal
	var err error
	var params hefloat.Parameters
	if params, err = hefloat.NewParametersFromLiteral(
		hefloat.ParametersLiteral{
			LogN:            12,
			LogQ:            []int{38, 32},
			LogP:            []int{39},
			LogDefaultScale: 32,
		}); err != nil {
		panic(err)
	}

	// Instantiate the key generator and then generate keys
	kgen := rlwe.NewKeyGenerator(params)
	sk := kgen.GenSecretKeyNew()
	pk := kgen.GenPublicKeyNew(sk)
	rlk := kgen.GenRelinearizationKeyNew(sk)
	evk := rlwe.NewMemEvaluationKeySet(rlk)

	// Randomly generate a value vector with element type `complex128` and size `Slots`,
	LogSlots := params.LogMaxSlots()
	Slots := 1 << LogSlots // `Slots` is N/2, i.e 2^12/2 = 2^11
	r := rand.New(rand.NewSource(0))
	values1 := make([]complex128, Slots)
	for i := 0; i < Slots; i++ {
		values1[i] = complex(2*r.Float64()-1, 2*r.Float64()-1) // both the real and imaginary part uniformly distributed in [-1, 1]
	}

	// Encode the value vector into a plaintext
	pt1 := hefloat.NewPlaintext(params, params.MaxLevel())
	ecd1 := hefloat.NewEncoder(params)
	if err = ecd1.Encode(values1, pt1); err != nil {
		panic(err)
	}

	// Encrypt the plaintext
	enc := rlwe.NewEncryptor(params, pk)
	ct1, err := enc.EncryptNew(pt1)
	if err != nil {
		panic(err)
	}

	// Randomly generate another value vector with element type `complex128` and size `Slots`
	values2 := make([]complex128, Slots)
	for i := 0; i < Slots; i++ {
		values2[i] = complex(2*r.Float64()-1, 2*r.Float64()-1)
	}

	// Encode the second value vector into a plaintext
	pt2 := hefloat.NewPlaintext(params, params.MaxLevel())
	ecd2 := hefloat.NewEncoder(hefloat.Parameters(params))
	if err = ecd2.Encode(values2, pt2); err != nil {
		panic(err)
	}

	// Encrypt the plaintext
	ct2, err := enc.EncryptNew(pt2)
	if err != nil {
		panic(err)
	}

	// Add the two complex value vectors
	want := make([]complex128, Slots)
	for i := 0; i < Slots; i++ {
		want[i] = values1[i] + values2[i]
	}

	// Perform the homomorphic addition
	eval := hefloat.NewEvaluator(params, evk)
	ct3, err := eval.AddNew(ct1, ct2)
	if err != nil {
		panic(err)
	}

	// Decrypt the result
	dec := rlwe.NewDecryptor(params, sk)
	pt3 := dec.DecryptNew(ct3)

	// Decode the result
	ecd3 := hefloat.NewEncoder(hefloat.Parameters(params))
	result := make([]complex128, Slots)
	if err = ecd3.Decode(pt3, result); err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		fmt.Printf("value with index %v in values1: %v\n", i, values1[i])
		fmt.Printf("value with index %v in values2: %v\n", i, values2[i])
		fmt.Printf("value with index %v in result:  %v\n", i, result[i])
		fmt.Printf("value with index %v in want:    %v\n", i, want[i])
		fmt.Println()
	}

	fmt.Println("...")
	fmt.Println()
	fmt.Printf("value with index %v in values1: %v\n", Slots-1, values1[Slots-1])
	fmt.Printf("value with index %v in values2: %v\n", Slots-1, values2[Slots-1])
	fmt.Printf("value with index %v in result:  %v\n", Slots-1, result[Slots-1])
	fmt.Printf("value with index %v in want:    %v\n", Slots-1, want[Slots-1])
}
