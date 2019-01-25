package main

func gen_xml_multi(UUID string) (string) {

	ans := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
	"<InitiateMultipartUploadResult xmlns=\"http://s3.amazonaws.com/doc/2006-03-01/\">" +
	"<UploadId>" + UUID + "</UploadId>" +
	"</InitiateMultipartUploadResult>"

	return ans

}
