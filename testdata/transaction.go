package testdata

const GetTransactionResponse = `{
  "Id": "67b9e140-b0f6-4c36-9b35-977041968185",
  "Status": 20,
  "Files": {
    "Contract.pdf": {
      "Links": [
        {
          "Rel": "file",
          "Type": "application/pdf",
          "Link": "https://api.signhost.com/api/transaction/67b9e140-b0f6-4c36-9b35-977041968185/file/Contract.pdf"
        }
      ],
      "DisplayName": "ArbeidsContract 2016"
    },
    "Algemenevoorwaarden": {
      "Links": [
        {
          "Rel": "file",
          "Type": "application/pdf",
          "Link": "https://api.signhost.com/api/transaction/67b9e140-b0f6-4c36-9b35-977041968185/file/Algemenevoorwaarden"
        }
      ],
      "DisplayName": "Algemene voorwaarden"
    }
  },
  "Seal": true,
  "Signers": [
    {
      "Id": "e0b39ec0-e0c6-45d9-bf0d-ae8cafbe2f48",
      "Expires": null,
      "Email": "user@example.com",
      "Mobile": "+31612345678",
      "Iban": null,
      "BSN": null,
      "RequireScribbleName": false,
      "RequireScribble": true,
      "RequireEmailVerification": true,
      "RequireSmsVerification": true,
      "RequireDigidVerification": false,
      "RequireSurfnetVerification": false,
      "SendSignRequest": true,
      "SendSignConfirmation": null,
      "SignRequestMessage": "Hello, could you please sign this document? Best regards, John Doe",
      "DaysToRemind": 15,
      "Language": "en-US",
      "ScribbleName": "John Doe",
      "ScribbleNameFixed": false,
      "Reference": "Client #123",
      "ReturnUrl": "http://signhost.com",
      "Activities": [
        {
          "Id": "91709d15-df2e-48a1-ac90-276f0360ce08",
          "Code": 103,
          "Activity": "Opened",
          "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00"
        },
        {
          "Id": "04adfda3-dd35-4f4d-af34-d2a08a4434f6",
          "Code": 203,
          "Activity": "Signed",
          "CreatedDateTime": "2016-06-15T23:38:04.1965465+02:00"
        }
      ],
      "RejectReason": null,
      "SignUrl": "http://ui.signhost.com/sign/d959e67b-acf8-4a49-8811-9b62f0b450af",
      "SignedDateTime": null,
      "RejectDateTime": null,
      "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "ModifiedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "Context": null
    }
  ],
  "Receivers": [
    {
      "Id": "df52316c-6671-4f39-9b9e-b524cc36ef93",
      "Name": "John Doe",
      "Email": "user@example.com",
      "Language": "en-US",
      "Message": "Hello, please find enclosed the digital signed document. Best regards, John Doe",
      "Reference": null,
      "Activities": null,
      "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "ModifiedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "Context": null
    }
  ],
  "Reference": "Contract #123",
  "PostbackUrl": "http://example.com/postback.php",
  "SignRequestMode": 2,
  "DaysToExpire": 30,
  "SendEmailNotifications": true,
  "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00",
  "ModifiedDateTime": "2016-06-15T23:33:04.1965465+02:00",
  "CanceledDateTime": null,
  "Context": null
}`

const CreateTransactionRequest = `{
	"Signers": [
    {
      "Email": "user@example.com",
      "SendSignRequest": true,
      "SignRequestMessage": "Hello, could you please sign this document? Best regards, John Doe",
      "DaysToRemind": 15,
      "Verifications": [
        {
          "Type": "iDeal"
        },
        {
          "Type": "Consent"
        }
      ]
    },
    {
      "Email": "anotheruser@example.com",
      "SendSignRequest": true,
      "SignRequestMessage": "Hello, could you please sign this document? Best regards, John Doe",
      "DaysToRemind": 15,
      "Verifications": [
        {
          "Type": "Consent"
        }
      ]
    }
    ],
    "SendEmailNotifications": true
}`

const CreateTransactionResponse = `{
  "Id": "67b9e140-b0f6-4c36-9b35-977041968185",
  "Status": 20,
  "Seal": true,
  "Signers": [
    {
      "Id": "e0b39ec0-e0c6-45d9-bf0d-ae8cafbe2f48",
      "Expires": null,
      "Email": "user@example.com",
      "Mobile": null,
      "Iban": null,
      "BSN": null,
      "RequireScribbleName": false,
      "RequireScribble": false,
      "RequireEmailVerification": false,
      "RequireSmsVerification": false,
      "RequireDigidVerification": false,
      "RequireSurfnetVerification": false,
      "Verifications": [
        {
          "Type": "iDeal"
		},
		{
			"Type": "Consent"
		}
      ],
      "SendSignRequest": true,
      "SendSignConfirmation": null,
      "SignRequestMessage": "Hello, could you please sign this document? Best regards, John Doe",
      "DaysToRemind": 15,
      "Language": "en-US",
      "ScribbleName": null,
      "ScribbleNameFixed": false,
      "Reference": "Client #123",
      "ReturnUrl": "http://signhost.com",
      "Activities": [],
      "RejectReason": null,
      "SignUrl": "http://ui.signhost.com/sign/d959e67b-acf8-4a49-8811-9b62f0b450af",
      "SignedDateTime": null,
      "RejectDateTime": null,
      "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "ModifiedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "Context": null
    },
    {
      "Id": "b9d0f613-985d-4a5a-8e79-a83f7b5d6b55",
      "Email": "anotheruser@example.com",
      "Mobile": null,
      "RequireScribble": false,
      "RequireSmsVerification": false,
      "SendSignRequest": true,
      "SignRequestMessage": "Hello, could you please sign this document? Best regards, John Doe",
      "DaysToRemind": 15,
      "ScribbleName": null,
	  "ScribbleNameFixed": false,
	  "Verifications": [
		  {
			  "Type": "Consent"
		  }
	  ]
    }
  ],
  "Receivers": [
    {
      "Id": "df52316c-6671-4f39-9b9e-b524cc36ef93",
      "Name": "John Doe",
      "Email": "user@example.com",
      "Language": "en-US",
      "Message": "Hello, please find enclosed the digital signed document. Best regards, John Doe",
      "Reference": null,
      "Activities": null,
      "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "ModifiedDateTime": "2016-06-15T23:33:04.1965465+02:00",
      "Context": null
    }
  ],
  "Reference": "Contract #123",
  "PostbackUrl": "http://example.com/postback.php",
  "SignRequestMode": 2,
  "DaysToExpire": 30,
  "SendEmailNotifications": true,
  "CreatedDateTime": "2016-06-15T23:33:04.1965465+02:00",
  "ModifiedDateTime": "2016-06-15T23:33:04.1965465+02:00",
  "CanceledDateTime": null,
  "Context": null
}`
