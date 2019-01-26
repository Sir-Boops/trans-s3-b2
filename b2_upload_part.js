const request = require('request')
const crypto = require('crypto')
const fs = require('fs')

module.exports = {
	b2_upload_part: function (auth, query, data, cb) {
		get_part_upload_url(auth, query.uploadId, function(upload_body){

			const options = {
				url: upload_body.uploadUrl,
				headers: {
					Authorization: upload_body.authorizationToken,
					'X-Bz-Part-Number': query.partNumber,
					'X-Bz-Content-Sha1': (crypto.createHash('sha1').update(data).digest('hex'))
				},
				body: data
			}

			request.post(options, function(err, res, body){
				// Write hash to a file
				fs.writeFileSync(body.fileId + '_' + body.partNumber, body.contentSha1)
				cb(res.statusCode)
			})
		})
	}
}

function get_part_upload_url(auth, fileid, cb) {
	const options = {
		url: (auth.api_url + '/b2api/v2/b2_get_upload_part_url'),
		headers: {
			Authorization: auth.auth
		},
		body: '{"fileId":"' + fileid + '"}'
	}

	request.post(options, function(err, res, body){
		cb(JSON.parse(body))
	})
}
