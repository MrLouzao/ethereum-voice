package main

// Contact information to detect who is the wallet address
type UserContact struct {
	Name			string
	WalletAddress	string
}

// Define an array of pre-defined contacts
// TODO: replace with database
var AppContacts = []UserContact {
	UserContact {
		Name: "Louzao",
		WalletAddress: "0xde95aff743b29b72885a27d03795632f6e741fdf",
	},
	UserContact {
		Name: "Pablo",
		WalletAddress: "0xd0CfE66448093dDA6cdC6525DB0C66BF1DD9c138",
	},
	UserContact {
		Name: "Rablo",
		WalletAddress: "0xd0CfE66448093dDA6cdC6525DB0C66BF1DD9c138",
	},
	UserContact {
		Name: "Alicia",
		WalletAddress: "0xCF46eC111492a5b00246eBBe260746A7E2E76558",
	},
	UserContact {
		Name: "maquina de caramelos",
		WalletAddress: "0xd0CfE66448093dDA6cdC6525DB0C66BF1DD9c138",
	},
}


// Obtain the contact wallet address if its contained in default AppContacts
// TODO replace this function with a call to database
func GetContactAddress(contactName string) string {
	var address string

	// Check if its in the array
	for i := 0; i < len(AppContacts); i++ {
		areContactsSamePerson := CompareStringIgnoringUpperAndAccents(AppContacts[i].Name, contactName)
		if areContactsSamePerson {
			address = AppContacts[i].WalletAddress
			break
		}
	}

	return address
}