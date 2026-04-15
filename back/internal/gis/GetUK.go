package gis

type UK struct {
	//Contacts UKContacts
	Phones []string `json:"phones"`
	Emails []string `json:"emails"`
	Name   string   `json:"name"`
	Addr   string   `json:"address"`
}

func GetUK(addr string) (*UK, error) {
	guid, regGuid, err := GetGuidByAddr(addr)
	if err != nil {
		return nil, err
	}

	house, err := GetIdByGuid(guid, regGuid)
	if err != nil {
		return nil, err
	}

	uk, err := GetUKIdById(house.Result.Id)
	if err != nil {
		return nil, err
	}

	ukInfo, err := GetUKContacts(uk.Id)
	if err != nil {
		return nil, err
	}

	return ukInfo, nil
}
