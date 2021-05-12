package testdata

const PutFileWithMetaRequest = `{
	"DisplayName": "Your personal contract",
	"Signers": {
		"SomeSignerId": {
			"FormSets": [ "SampleFormset" ]
		}
	},
	"FormSets": {
		"SampleFormset": {
			"AddressLine1": {
				"Type": "SingleLine",
				"Location": {
					"Search": "Address line 1",
					"Left": 100
				}
			},
			"AddressLine2": {
				"Type": "SingleLine",
				"Location": {
					"Search": "Address line 2",
					"Left": 100
				}
			},
			"SignatureOne": {
				"Type": "Signature",
				"Location": {
					"Right": 10,
					"Top": 10,
					"PageNumber": 1,
					"Width": 140,
					"Height": 70
				}
			}
		},
		"SecondSigner": {
			"Signature-2": {
				"Type": "Signature",
				"Location": {
					"PageNumber": 2,
					"Width": 140,
					"Height": 70
				}

			}
		}
	}
}`

const PutFileWithMetaResponse =	`{
	"status": "ok"
}`

