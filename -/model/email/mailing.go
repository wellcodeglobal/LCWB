package email

import (
	"fmt"
	baseurl "github.com/kevinrizkhy/Internal/config"
	config "github.com/kevinrizkhy/Internal/config"
	database "github.com/kevinrizkhy/Internal/db"
	"github.com/kevinrizkhy/Internal/lib/utility"
	"github.com/scorredoira/email"
	"io/ioutil"
	"log"
	//"net/http"
	"net/mail"
	"net/smtp"
	"strings"
)

var count int = 0
var (
	role         string
	first_name   string
	last_name    string
	user_email   string
	company_name string
	address_work string
	contact_work string
	phone        string
	ktp          string
	gender       string
	reason       string
	joindate     string
	birthday     string
	position     string
	district     string
	company_file string
)

func sendMail(subj, mess, receipient, type_email string) bool {
	var rows map[int](map[string]string)
	var err error
	if type_email == "Campaign" {
		rows, err = database.ExecuteQuery("SELECT name, email, password, smtp, port FROM administrative_email WHERE name LIKE '%sales%'")
	} else {
		rows, err = database.ExecuteQuery("SELECT name, email, password, smtp, port FROM administrative_email WHERE name LIKE '%system%'")
	}
	if len(rows) > 0 && err == nil {
		for i := 0; i < len(rows); i++ {
			name := rows[i]["name"]
			user := rows[i]["email"]
			password := rows[i]["password"]
			host := rows[i]["smtp"]
			port := rows[i]["port"]
			m := email.NewHTMLMessage(subj, mess)
			m.From = mail.Address{Name: "PT. Sinar Inti Panel Indonesia", Address: name}
			m.To = []string{receipient}
			auth := smtp.PlainAuth("", user, password, host)
			err := email.Send(host+":"+port, auth, m)
			if err != nil {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func SendMailAttachment(subj, mess, receipient, type_email, file string) bool {
	var rows map[int](map[string]string)
	var err error
	//Ada yang diubah
	//type_email := "Campaign"
	if type_email == "Campaign" {
		rows, err = database.ExecuteQuery("SELECT name, email, password, smtp, port FROM administrative_email WHERE name LIKE '%sales%'")
	} else {
		rows, err = database.ExecuteQuery("SELECT name, email, password, smtp, port FROM administrative_email WHERE name LIKE '%system%'")
	}
	if len(rows) > 0 && err == nil {
		for i := 0; i < len(rows); i++ {
			m := email.NewHTMLMessage(subj, mess)
			m.From = mail.Address{Name: "PT. Sinar Inti Panel Indonesia", Address: rows[i]["name"]}
			m.To = []string{receipient}
			fmt.Println(file)
			m.Attach(file)
			auth := smtp.PlainAuth("", rows[i]["email"], rows[i]["password"], rows[i]["smtp"])
			err := email.Send(rows[i]["smtp"]+":"+rows[i]["port"], auth, m)
			if err != nil {
				return false
			} else {
				return true
			}
		}
	}
	return false
}

func SendContactUsEmail(receipient string, token string, file string) bool { //, data ContentEmail) {
	url := baseurl.Base_URL + "verificationdata?token=" + token
	body := GetHTMLConfirmation(url)
	fileupload := "upload/" + file
	subject := "Konfirmasi Distributor PT. Sinar Inti Panel"
	typeemail := "Campaign"
	return SendMailAttachment(subject, body, receipient, typeemail, fileupload)
}

// func SendAtt(w http.ResponseWriter, r *http.Request) {
// 	url := baseurl.Base_URL + "#"
// 	body := GetHTMLConfirmation(url)
// 	fileupload := "upload/receipt/3_Alphabeth Corp_1514246400.pdf"
// 	subject := "Konfirmasi Distributor PT. Sinar Inti Panel"
// 	receipient := "fadhilahsan2104@gmail.com"
// 	SendMailAttachment(subject, body, receipient, fileupload)
// }

func GetHTMLConfirmationEmail(token string) (string, string, string) {
	url := baseurl.Base_URL + "verificationdata?token=" + token
	body := GetHTMLConfirmation(url)
	//fileupload := "upload/" + file
	subject := "Konfirmasi Distributor PT. Sinar Inti Panel"
	typeemail := "Confirmation"
	return body, subject, typeemail
}

func ResendEmailConfirmation(email string) bool {
	rows, err := database.ExecuteQuery("SELECT status FROM users WHERE email=$1", email)
	if len(rows) > 0 && err == nil {
		status := rows[0]["status"]
		if status == "Need Confirmation" {
			newToken := utility.GenerateToken(email)
			_, err := database.ExecuteQuery("update users set registration_token=$1 where email=$2 AND status=$3", newToken, email, status)
			if err == nil {
				mess, subj, _ := GetHTMLConfirmationEmail(newToken)
				_, err2 := database.ExecuteQuery("INSERT INTO public.cron_email( target, subject, message, type, reference) VALUES ($1,$2,$3,$4,$5)", email, subj, mess, "Resend Confirmation Email", 0)
				if err2 != nil {
					fmt.Println(err2)
					return false
				} else {
					return true
				}
			} else {
				fmt.Println(err)
				return false
			}
		} else {
			fmt.Println(err)
			return false
		}
	} else {
		fmt.Println(err)
		return false
	}
}

func GetHTMLEmailAgreement(tokenapproval string) (string, string, string) {
	subject := "Persetujuan Perjanjian Kerja Sama Distributor PT. Sinar Inti Panel"
	url := baseurl.Base_URL + "agreementdata?token=" + tokenapproval
	body := GetHTMLAgrrement(url)
	typeemail := "Agreement"
	return body, subject, typeemail
	//return sendMail(subject, body, email)
}

func GetHTMLForgotPassword(tokenapproval string) string {
	//subject := "Persetujuan Perjanjian Kerja Sama Distributor PT. Sinar Inti Panel"
	url := baseurl.Base_URL + "resetpassword?token=" + tokenapproval
	body := GetForgotPasswordHTML(url)
	//typeemail := "Agreement"
	return body
	//return sendMail(subject, body, email)
}

func GetHTMLEmailApprovalUser(usertoken string) (string, string, string, []string) {
	var targets []string
	check := utility.CheckVerified(usertoken)
	if check {
		rows, err := database.ExecuteQuery("SELECT first_name,last_name,email,company_name,address_work,contact_work,phone,ktp,birthday,gender,approval_token,company_file,position,region FROM users WHERE registration_token=$1 AND verified=$2", usertoken, "false")
		if len(rows) > 0 && err == nil {
			first_name = rows[0]["first_name"]
			last_name = rows[0]["last_name"]
			user_email = rows[0]["email"]
			company_name = rows[0]["company_name"]
			address_work = rows[0]["address_work"]
			contact_work = rows[0]["contact_work"]
			phone = rows[0]["phone"]
			ktp = rows[0]["ktp"]
			birthday = rows[0]["birthday"]
			gender = rows[0]["gender"]
			token_approval := rows[0]["approval_token"]
			company_file = rows[0]["company_file"]
			position = rows[0]["position"]
			district = rows[0]["region"]
			rows1, err1 := database.ExecuteQuery("SELECT id_region,nama FROM region WHERE id_region=$1", district)
			if len(rows1) > 0 && err1 == nil {
				birthday = birthday[:10]
				id_region := rows1[0]["id_region"]
				nama_region := rows1[0]["nama"]
				rows2, err3 := database.ExecuteQuery("SELECT nama FROM city WHERE id_region=$1", id_region)
				if len(rows2) > 0 && err3 == nil {
					district = nama_region + "("
					for i := 0; i < len(rows2); i++ {
						district = district + " " + rows2[i]["nama"] + ","
					}
					district = strings.TrimRight(district, ",")
					district = district + ")"
					subject := "Persetujuan Distributor PT. Sinar Inti Panel"
					url := baseurl.Base_URL + "approve?token=" + token_approval
					msg := GetHTMLApprove(first_name, last_name, user_email, position, company_name, district, address_work, contact_work, phone, ktp, gender, birthday, url)
					typeemail := "Approval User"
					targetList, err2 := database.ExecuteQuery("select email from users left join groups on users.id_group=groups.id_groups WHERE groups.nama=$1;", "Admin")
					if len(targetList) > 0 && err2 == nil {
						for i := 0; i < len(targetList); i++ {
							targets = append(targets, targetList[0]["email"])
						}
					} else {
						if err2 != nil {
							fmt.Println("select email -->", err)
						}
					}
					return msg, subject, typeemail, targets
				} else {
					fmt.Println(err)
					return "", "", "", targets
				}
			} else {
				fmt.Println(err)
				return "", "", "", targets
			}
		} else {
			fmt.Println(err)
			return "", "", "", targets
		}
	} else {
		return "", "", "", targets
	}
}

func GetHTMLPaidoffInvoiceFinance(url, invoicenumber, company_name string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailPaidoffInvoiceFinance.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		//temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Company_name}}", company_name, -1)
		temp = strings.Replace(temp, "{{.Invoice_number}}", invoicenumber, -1)
		return temp
	}
}

func GetHTMLPaidoffInvoiceDistributor(url, invoicenumber string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailPaidoffInvoiceDsitributor.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		//temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Invoice_number}}", invoicenumber, -1)
		return temp
	}
}

func GetHTMLRejectPaymentDistributor(url, invoicenumber string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailRejectClaimPaymentDsitributor.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		//temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Invoice_number}}", invoicenumber, -1)
		return temp
	}
}

func GetHTMLApproveClaimDistributor(url, invoicenumber string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailApproveClaimPaymentDsitributor.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Invoice_number}}", invoicenumber, -1)
		return temp
	}
}

func GetHTMLClaimFinance(url string, company_name string, invoice string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailClaimPaymentFinance.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Comapany_name}}", company_name, -1)
		temp = strings.Replace(temp, "{{.Invoice_number}}", invoice, -1)
		return temp
	}
}

func GetHTMLNewOrderSales(url string, company_name string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailOrderTemplateSales.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Comapany_name}}", company_name, -1)
		return temp
	}
}

func GetHTMLNewOrderDistributor() string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailOrderTemplateDistributor.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		return temp
	}
}

func GetHTMLOfferDistributor1(url string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailOfferTemplateDistributor1.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		return temp
	}
}

func GetHTMLOfferDistributor2(url string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailOfferTemplateDistributor2.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		return temp
	}
}

func GetHTMLEmailApprovalEditUser(iduser string) (string, string, string, []string) {
	var targets []string
	rows, err := database.ExecuteQuery("SELECT company_name FROM users WHERE id_user=$1 AND verified=$2", iduser, "true")
	if len(rows) > 0 && err == nil {
		company_name = rows[0]["company_name"]
		rows3, err5 := database.ExecuteQuery("select token from data_change where id_user=$1 order by create_date DESC", iduser)
		if len(rows3) > 0 && err5 == nil {
			token := rows3[0]["token"]
			subject := "Persetujuan Perubahan Data Distributor PT. Sinar Inti Panel"
			url := baseurl.Base_URL + "approvaledituser?token=" + token
			msg := GetHTMLEditProfileAdmin(url, company_name)
			typeemail := "Edit User"
			targets := strings.Split(config.ADMIN, ",")
			return msg, subject, typeemail, targets
		} else {
			fmt.Println(err)
			return "", "", "", targets
		}
	} else {
		fmt.Println(err)
		return "", "", "", targets
	}

}

func GetHTMLRejectEditProfileDistributor() string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2_RejectEditProfile.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		return temp
	}
}

func GetHTMLApproveEditProfileDistributor() string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2_ApproveEditProfile.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		return temp
	}
}

func GetHTMLEditProfileAdmin(url string, company_name string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailtemplate2_editprofileadmin.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		temp = strings.Replace(temp, "{{.Comapany_name}}", company_name, -1)
		return temp
	}
}

func GetHTMLEditProfileDistributor() string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailtemplate2_editprofiledistributor.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		return temp
	}
}

func GetHTMLApprove(first_name, last_name, email, jabatan, nama_perusahaan, wilayah, alamat_perusahaan, telepon_saat, nomor_handphone, nomot_ktp, jenis_kelamin, tanggal_lahir, url string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		name := first_name + " " + last_name
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.nama}}", name, -1)
		temp = strings.Replace(temp, "{{.email}}", email, -1)
		temp = strings.Replace(temp, "{{.jabatan}}", jabatan, -1)
		temp = strings.Replace(temp, "{{.nama_perusahaan}}", nama_perusahaan, -1)
		temp = strings.Replace(temp, "{{.alamat_perusahaan}}", alamat_perusahaan, -1)
		temp = strings.Replace(temp, "{{.telepon_saat}}", telepon_saat, -1)
		temp = strings.Replace(temp, "{{.nomor_handphone}}", nomor_handphone, -1)
		temp = strings.Replace(temp, "{{.nomor_ktp}}", nomot_ktp, -1)
		temp = strings.Replace(temp, "{{.tanggal_lahir}}", tanggal_lahir, -1)
		if jenis_kelamin == "M" {
			temp = strings.Replace(temp, "{{.jenis_kelamin}}", "Laki-laki", -1)
		} else {
			temp = strings.Replace(temp, "{{.jenis_kelamin}}", "Perempuan", -1)
		}
		temp = strings.Replace(temp, "{{.region_distribusi}}", wilayah, -1)
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		return temp
	}
}

func GetHTMLConfirmation(url string) string {

	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2_client.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		return temp
	}
}

func GetHTMLApproveClient() string {

	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2_client2.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		return temp
	}

}

func GetHTMLAgrrement(url string) string {

	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2_client3.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		return temp
	}
}

func SendCronEmail() bool {
	rows, err := database.ExecuteQuery("SELECT id,target,subject, message,type,attachment FROM cron_email where date_sent is null")
	if len(rows) > 0 && err == nil {
		for i := 0; i < len(rows); i++ {
			message := rows[i]["message"]
			message = strings.Replace(message, "[emailID]", rows[i]["id"], 1)
			attachments := rows[i]["attachment"]
			if attachments == "#" {
				if sendMail(rows[i]["subject"], rows[i]["message"], rows[i]["target"], rows[i]["type_email"]) {
					_, err := database.ExecuteQuery("update cron_email set date_sent=NOW() where id=$1", rows[i]["id"])
					if err == nil {
						fmt.Println("Send mail to " + rows[i]["target"] + " success!")
					}
				} else {
					fmt.Println("Exit IF - sendEmail : False")
				}
			} else {
				if SendMailAttachment(rows[i]["subject"], rows[i]["message"], rows[i]["target"], rows[i]["type_email"], attachments) {
					_, err := database.ExecuteQuery("update cron_email set date_sent=NOW() where id=$1", rows[i]["id"])
					if err == nil {
						fmt.Println("Send mail to " + rows[i]["target"] + " success!")
					}
				} else {
					fmt.Println("Exit IF - sendEmail : False")
				}
			}
		}
	} else {
		return false
	}
	return true
}

func GetForgotPasswordHTML(url string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailTemplate2_forgotpassword.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.url}}", url, -1)
		return temp
	}
}

func GetHTMLClaimDistributor(invoicenumber string) string {
	content, err := ioutil.ReadFile("app/views/dynamic/template/emailClaimPaymentDistributor.html")
	if err != nil {
		log.Println(err)
		return ""
	} else {
		temp := string(content[:])
		temp = strings.Replace(temp, "{{.Number_Invoice}}", invoicenumber, -1)
		return temp
	}
}
