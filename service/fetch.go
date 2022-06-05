package service

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch requests HTTP body by API URL
func Fetch(fullUrl string) (rs []byte, err error) {
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Set("cookie", "memclid=db2a3ed5-d955-4bc0-b529-a8ec72cabb52; nfvdid=BQFmAAEBEHZCc685X3sWljCjtkjd8nFA02b6%2BXIIH0jutOm6zMKhq8%2BrYMFyGbDeezNdU8m3EOwK7CGd2DIpBb1S2UivvKb1TQf%2B2oxcIIQtCELIhfMwhw%3D%3D; flwssn=06583178-8540-43c8-a7df-932c7fb168b9; SecureNetflixId=v%3D2%26mac%3DAQEAEQABABTl8T2Tv19WIYKkDy1WaTTzCkq_qUHx_KA.%26dt%3D1654350938123; NetflixId=v%3D2%26ct%3DBQAOAAEBEE2aODjOn0Yqe9FSdSFpl0WBAEUa2dfwmBDyYHt_u0F1-xTTQQWeLQe5YnNgKgKP-dXE5f2cPRRWlVP_Qumg8DbKvx06t-PQsX_pEtVz_UKV2Vv-FRfx8vZE8SwnJPpcR9IpXQjMx_coTF3UOc00LP5JwCm42dQxsiVZYuoDsJrPX04-HdMTJQmi0v9UGRsbQ9ZEkleP31pDIZ-UOILwlMjOhpDQXPNJdQUSWDv5gIGTx_mHwCi5eyNFj20OD_4QNK5fbUZGiZBbCkcjVVefwjbvNiykh0Z1YKuWJKpHg_58SdvIjXU9042R6mzInHVFYfolUpthFkF9GyLNwShFBYuwuzGNHFrVUk1cKzmHrevnOc4.%26bt%3Ddev%26mac%3DAQEAEAABABSl-d5Ytiy59r8-oQ4yynP9EdCn0koUN8U.; cL=1654351687518%7C165435024421014994%7C165435024429887127%7C%7C32%7Cundefined; OptanonConsent=isIABGlobal=false&datestamp=Sat+Jun+04+2022+22%3A08%3A14+GMT%2B0800+(%E4%B8%AD%E5%9B%BD%E6%A0%87%E5%87%86%E6%97%B6%E9%97%B4)&version=6.6.0&consentId=087f4c7f-139f-458c-8172-a096179033ae&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1&hosts=H12%3A1%2CH13%3A1%2CH51%3A1%2CH45%3A1%2CH46%3A1%2CH48%3A1%2CH49%3A1&AwaitingReconsent=false")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	log.Printf("%+v", resp)

	rs, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	//body长度太短,可以判断没有访问成功
	if len(rs) < MinBodyLength {
		return rs, errors.New("Not Available ")
	}

	return rs, nil
}
