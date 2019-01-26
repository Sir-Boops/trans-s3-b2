const request = require('request')
const crypto = require('crypto')

module.exports = {
	b2_upload: function(auth, path, body, cb) {

		// First we need to get the upload URL
		get_upload_url(auth, function(upload_path) {

			// Now upload the file
			const options = {
				url: upload_path.uploadUrl,
				headers: {
					Authorization: upload_path.authorizationToken,
					'X-Bz-File-Name': path.replace('/' + auth.bucket_name + '/', ''),
					'Content-Type': 'b2/x-auto',
					'X-Bz-Content-Sha1': (crypto.createHash('sha1').update(body).digest('hex'))
				},
				body: body
			}

			request.post(options, function(err, res, body){
				cb(res.statusCode)
			})

		})
	}
}

function get_upload_url(auth, cb) {

	const options = {
		url: (auth.api_url + "/b2api/v2/b2_get_upload_url"),
		headers: {
			Authorization: auth.auth
		},
		body: '{ "bucketId": "' + auth.bucket_id + '" }'
	}

	request.post(options, function(err, res, body) {
		cb(JSON.parse(body))
	})
}
