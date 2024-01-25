package peda

import (
	"time"
)

type Pesan struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data,omitempty" bson:"data,omitempty"`
	Token   string      `json:"token,omitempty" bson:"token,omitempty"`
}
type CredentialUser struct {
	Status bool `json:"status" bson:"status"`
	Data   struct {
		No_whatsapp string `json:"no_whatsapp" bson:"no_whatsapp"`
		Username    string `json:"username" bson:"username"`
		Role        string `json:"role" bson:"role"`
	} `json:"data" bson:"data"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type User struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password,omitempty"`
	Role        string `json:"role,omitempty" bson:"role,omitempty"`
	Token       string `json:"token,omitempty" bson:"token,omitempty"`
	Private     string `json:"private,omitempty" bson:"private,omitempty"`
	Publick     string `json:"publick,omitempty" bson:"publick,omitempty"`
	No_whatsapp string `json:"no_whatsapp,omitempty" bson:"no_whatsapp,omitempty"`
}

type Dosen struct {
	Pendidikan_dosen  string `json:pendidikan_dosen bson:"pendidikan_dosen"`
	Kuriulum_dosen    string `json:kuriulum_dosen bson:"kuriulum_dosen"`
	Penelitian_dosen  string `json:penelitian_dosen bson:"penelitian_dosen"`
	Gelar_dosen       string `json:gelar_dosen bson:"gelar_dosen"`
	Lembaga_dosen     string `json:lembaga_dosen bson:"lembaga_dosen"`
	Kemampuan_dosen   string `json:kemampuan_dosen bson:"kemampuan_dosen"`
	Penghargaan_dosen string `json:penghargaan_dosen bson:"penghargaan_dosen"`
}

type UserToken struct {
	Username User `json:"username" bson:"username"`
}

type Payload struct {
	No_whatsapp string    `json:"no_whatsapp"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	Exp         time.Time `json:"exp"`
	Iat         time.Time `json:"iat"`
	Nbf         time.Time `json:"nbf"`
}

type Credential struct {
	Status   bool        `json:"status" bson:"status"`
	Token    string      `json:"token,omitempty" bson:"token,omitempty"`
	Message  string      `json:"message,omitempty" bson:"message,omitempty"`
	Username string      `json:"username,omitempty" bson:"username,omitempty"`
	Data     interface{} `json:"data,omitempty" bson:"data,omitempty"`
}

type Response struct {
	Status  bool        `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
}

type Backend struct {
	Npm                int    `json:npm bson:"npm"`
	Nama               string `json:nama bson:"nama"`
	NamaDosen          string `json:namadosen bson:"namadosen"`
	Autentikasitoken   string `json:autentikasitoken bson:"autentikasitoken"`
	Packagesendiri     string `json:packagesendiri bson:"packagesendiri"`
	Endpointgcfjakarta string `json:endpointgcfjakarta bson:"endpointgcfjakarta"`
	Integrasiwamyid    string `json:integrasiwamyid bson:"integrasiwamyid"`
	Status             bool   `json:status bson:"status"`
}
type Frontend struct {
	Npm            int    `json:npm bson:"npm"`
	Nama           string `json:nama bson:"nama"`
	NamaDosen      string `json:namadosen bson:"namadosen"`
	Rilisjs        string `json:rilisjs bson:"rilisjs"`
	Pemanggilanjs  string `json:pemanggilanjs bson:"pemanggilanjs"`
	Kelengkapancss string `json:kelengkapancss bson:"kelengkapancss"`
	CustomDomain   string `json:customdomain bson:"customdomain"`
	Status         bool   `json:status bson:"status"`
}

type FormInputAll struct {
	Nama_dosen  string        `json:nama_dosen bson:"nama_dosen"`
	Nik         string        `json:nik bson:"nik"`
	Dosen       []Dosen       `json:dosen bson:"dosen"`
	Suratall    []SuratKerja  `json:suratkerja bson:"suratkerja"`
	Sertificate []Sertificate `json:sertificate bson:"sertificate"`
	Akademis    []Akademis    `json:akademis bson:"akademis"`
	Status      bool          `json:status bson:"status"`
}

type SuratKerja struct {
	Penawaran_kerja     string `json:penawaran_kerja bson:"penawaran_kerja"`
	Perjanjian_kerja    string `json:perjanjian_kerja bson:"perjanjian_kerja"`
	Pemberhentian_kerja string `json:pemberhentian_kerja bson:"pemberhentian_kerja"`
	Keterangan_kerja    string `json:keterangan_kerja bson:"keterangan_kerja"`
	Kuasa_kerja         string `json:kuasa_kerja bson:"kuasa_kerja"`
}

type Sertificate struct {
	Judul_sertifikat              string `json:judul_sertifikat bson:"judul_sertifikat"`
	Pemberi_sertifikat            string `json:pemberi_sertifikat bson:"pemberi_sertifikat"`
	Penerima_sertifikat           string `json:penerima_sertifikat bson:"penerima_sertifikat"`
	Tujuan_sertifikat             string `json:tujuan_sertifikat bson:"tujuan_sertifikat"`
	Tanggal_penerbitan_sertifikat string `json:tanggal_penerbitan_sertifikat bson:"tanggal_penerbitan_sertifikat"`
	Cap_sertifikat                string `json:cap_sertifikat bson:"cap_sertifikat"`
	Nomo_sertifikat               string `json:nomo_sertifikat bson:"nomo_sertifikat"`
	Info_sertifikat               string `json:info_sertifikat bson:"info_sertifikat"`
	Logo_sertifikat               string `json:logo_sertifikat bson:"logo_sertifikat"`
}

type Akademis struct {
	Nama_dosen        string `json:nama_dosen bson:"nama_dosen"`
	Pendidikan_dosen  string `json:pendidikan_dosen bson:"pendidikan_dosen"`
	Kuriulum_dosen    string `json:kuriulum_dosen bson:"kuriulum_dosen"`
	Penelitian_dosen  string `json:penelitian_dosen bson:"penelitian_dosen"`
	Gelar_dosen       string `json:gelar_dosen bson:"gelar_dosen"`
	Lembaga_dosen     string `json:lembaga_dosen bson:"lembaga_dosen"`
	Kemampuan_dosen   string `json:kemampuan_dosen bson:"kemampuan_dosen"`
	Penghargaan_dosen string `json:penghargaan_dosen bson:"penghargaan_dosen"`
}

type DocumentInput struct {
	filetypes    string `json:filetypes bson:"filetypes"`
	Documentfile string `json:documents bson:"documents"`
}

type Document struct {
	Encrypted_Docs Backend `json:document bson:"document"`
}
