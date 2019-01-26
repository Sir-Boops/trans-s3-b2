const request = require('request');

module.exports = {
	b2_start_large_upload: function (auth, path, cb) {

		options = {
			url: (auth.api_url + '/b2api/v2/b2_start_large_file'),
			headers: {
				Authorization: auth.auth
			},
			body: '{"fileName":"' + path.replace('/' + auth.bucket_name + '/', '') + '","bucketId":"' + auth.bucket_id + '","contentType":"b2/x-auto"}'
		}

		request.post(options, function(err, res, body){
			cb(res.statusCode, JSON.parse(body))
		})
	}
}
