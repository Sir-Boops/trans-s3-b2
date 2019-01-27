const request = require('request')

module.exports = {
	b2_finish_large_file: function(auth, uploadId, db, cb) {
		gen_list(db, uploadId, function(list){

			const options = {
				url: (auth.api_url + "/b2api/v2/b2_finish_large_file"),
				headers: {
					Authorization: auth.auth
				},
				body: '{"fileId":"' + uploadId + '", ' + list + '}'
			}

			request.post(options, function(err, res, body){
				console.log(body)
				cb(res.statusCode)
			})
		})
	}
}

function gen_list(db, uploadId, cb) {
	db.all('SELECT hash FROM hashes WHERE id LIKE ? ORDER BY CAST(part AS INTEGER)', [uploadId], function(err, rows){
		ans = '['
		var i;
		for(i = 0; i < rows.length; i++){
			if((i + 1) < rows.length){
				ans += ('"' + rows[i].hash + '",')
			}
			if((i + 1) >= rows.length){
				ans += ('"' + rows[i].hash + '"')
				ans += ']'
				cb(ans)
			}
		}
		if(rows.length == 0){
			cb('[]')
		}
	})
}
