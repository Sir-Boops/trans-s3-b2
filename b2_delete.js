const request = require('request');

module.exports = {
	b2_delete: function(auth, path, cb) {

		// Get file ID
		hide_file(auth, path, function(code, hide_body){

			if(code == 200) {

				const options = {
					url: (auth.api_url + '/b2api/v2/b2_delete_file_version'),
					headers: {
						Authorization: auth.auth
					},
					body: '{"fileName":"' + hide_body.fileName + '","fileId":"' + hide_body.fileId + '"}'
				}

				request.post(options, function(err, res, body){
					cb(res.statusCode)
				})

			} else {
				if(code != 404){
					cb(code)
				}
			}
		})
	}
}

function hide_file(auth, path, cb) {

	const options = {
		url: (auth.api_url + '/b2api/v2/b2_hide_file'),
		headers: {
			Authorization: auth.auth
		},
		body: '{"bucketId":"' + auth.bucket_id + '", "fileName":"' + path.replace('/' + auth.bucket_name + '/', '') + '"}'
	}

	request.post(options, function(err, res, body){
		cb(res.statusCode, JSON.parse(body))
	})

}
