package util

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
}

type Meta struct {
	Page       uint32 `json:"page"`
	Limit      uint32 `json:"limit"`
	TotalRows  uint32 `json:"total_rows"`
	TotalPages uint32 `json:"total_pages"`
}

const (
	MESSAGE_OK_CREATE              = "Berhasil memasukkan data."
	MESSAGE_OK_DELETE              = "Berhasil menghapus data."
	MESSAGE_OK_UPDATE              = "Berhasil memperbaharui data."
	MESSAGE_OK_READ                = "Berhasil menampilkan data."
	MESSAGE_OK_IMPORT              = "Berhasil mengimport data."
	MESSAGE_OK_EXPORT              = "Berhasil mengexport data."
	MESSAGE_OK_UPLOAD              = "Berhasil mengunggah data"
	MESSAGE_OK_LOGIN               = "Berhasil masuk."
	MESSAGE_OK_REFRESH             = "Berhasil membuat ulang token."
	MESSAGE_OK_LOGOUT              = "Berhasil keluar."
	MESSAGE_OK_REQUEST_TOKEN_RESET = "Berhasil meminta token."
	MESSAGE_OK_CHANGE_PASSWORD     = "Berhasil memperbaharui password."
	MESSAGE_FAILED_CREATE          = "Gagal memasukkan data data."
	MESSAGE_FAILED_DELETE          = "Gagal menghapus data."
	MESSAGE_FAILED_UPDATE          = "Gagal memperbaharui data."
	MESSAGE_FAILED_READ            = "Gagal menampilkan data."
	MESSAGE_FAILED_IMPORT          = "Gagal mengimport data."
	MESSAGE_FAILED_EXPORT          = "Gagal mengexport data."
	MESSAGE_BAD_REQUEST            = "Permintaan bermasalah."
	MESSAGE_BAD_SYSTEM             = "Server bermasalah."
	MESSAGE_UNAUTHORIZED           = "Unauthorized."
	MESSAGE_CONFLICT               = "Data sudah ada."
	MESSAGE_NOT_FOUND              = "Data tidak ada."
	MESSAGE_FAILED_LOGIN           = "Username atau password tidak cocok."
)
