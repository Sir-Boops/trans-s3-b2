const request = require('request');

module.exports = {
	b2_auth: function(id, key, cb) {

		const options = {
			url: 'https://api.backblazeb2.com/b2api/v2/b2_authorize_account',
			headers: {
				'Authorization': ("Basic " + Buffer.from(id + ":" + key).toString('base64'))
			}
		}

		request.get(options, function (err, res, body) {

			body = JSON.parse(body)

			const ans = {
				bucket_id: body.allowed.bucketId,
				bucket_name: body.allowed.bucketName,
				api_url: body.apiUrl,
				auth: body.authorizationToken,
				download_url: body.downloadUrl
			}

			cb(ans)
		})
	}
}
