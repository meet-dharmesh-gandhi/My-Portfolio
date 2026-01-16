package ast

import (
	"testing"
)

func TestAST(t *testing.T) {
	var arrMin float64 = 1
	var arrMax float64 = 10
	var stringLen float64 = 3
	var mapMin float64 = 6
	var mapMax float64 = 16
	var node TypeNode = &Array{
		Length: NumberRange{
			Min: &arrMin,
			Max: &arrMax,
		},
		ElementType: CompoundType{
			Options: []TypeNode{
				String{
					Length: &NumberRange{
						Min: &stringLen,
						Max: &stringLen,
					},
				},
				Int{
					Number: &NumberRange{
						Min: &arrMin,
					},
				},
				Map{
					Length: NumberRange{
						Min: &mapMin,
						Max: &mapMax,
					},
					KeyType: CompoundType{
						Options: []TypeNode{
							String{
								Length: &NumberRange{
									Min: &stringLen,
									Max: &stringLen,
								},
							},
							Int{
								Number: &NumberRange{
									Min: &arrMin,
								},
							},
						},
					},
					ValueType: CompoundType{
						Options: []TypeNode{
							LiteralArray{
								ElementTypes: []CompoundType{
									{
										Options: []TypeNode{
											String{
												Length: &NumberRange{
													Min: &stringLen,
													Max: &stringLen,
												},
											},
											Int{
												Number: &NumberRange{
													Min: &arrMin,
												},
											},
										},
									},
									{
										Options: []TypeNode{
											String{
												Length: &NumberRange{
													Min: &stringLen,
													Max: &stringLen,
												},
											},
										},
									},
								},
							},
							LiteralMap{
								Elements: []MapEntry{
									{
										Key: "abc",
										ValueType: CompoundType{
											Options: []TypeNode{
												Int{
													Number: &NumberRange{
														Min: &arrMin,
													},
												},
											},
										},
									},
									{
										Key: "def",
										ValueType: CompoundType{
											Options: []TypeNode{
												Int{
													Number: &NumberRange{
														Min: &arrMin,
														Max: &arrMax,
													},
												},
											},
										},
									},
								},
							},
						},
					},
					ExplicitTypes: []MapOverrides{
						{
							ExplicitKey: "some key",
							ValueType: CompoundType{
								Options: []TypeNode{
									String{
										Length: &NumberRange{
											Min: &stringLen,
											Max: &stringLen,
										},
									},
									Int{
										Number: &NumberRange{
											Min: &arrMin,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		ExplicitTypes: []ArrayOverrides{
			{
				Index: NumberRange{},
				ElementType: CompoundType{
					Options: []TypeNode{
						String{
							Length: &NumberRange{
								Min: &stringLen,
								Max: &stringLen,
							},
						},
					},
				},
			},
		},
	}

	PrintType(node, 0)
}
