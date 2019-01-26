const express = require('express')
const getRawBody = require('raw-body')
const sqlite3 = require('sqlite3')
const b2_auth = require('./b2_auth')
const b2_upload = require('./b2_upload')
const b2_delete = require('./b2_delete')
const b2_start_large_upload = require('./b2_start_large_upload')
const b2_upload_part = require('./b2_upload_part')

const app = express();
let db = new sqlite3.Database(':memory:');

// Create the multi-part in memory DB
console.log('Starting DB')
db.run('CREATE TABLE hashes (id TEXT, hash TEXT)', function(err){
	console.log('DB started')
	// Auth to b2
	console.log("Logging into Backblaze")
	b2_auth.b2_auth(process.argv[2], process.argv[3], function(auth) {
		console.log("Logged into Backblaze")

		console.log('Starting server')

		// Handle PUT requests
		app.put("/*", function(req, res){
			console.log(res.query)
			getRawBody(req, function(err, data){
				if(data !== undefined) {
					if(req.query.uploadId !== undefined && req.query.partNumber !== undefined) {
						res.status(200)
						res.send()
						//Large part file upload
						b2_upload_part.b2_upload_part(auth, req.query, data, db, function(code){
							if(code == 200){
								console.log('Uploaded a part of a large file!')
							}
						})
					} else {
						// Normal small file upload
						b2_upload.b2_upload(auth, req.path, data, function(code) {
							if(code == 200) {
								console.log('Uploaded file: ' + req.path)
							} else {
								console.log('Non 200 status when uploading: ' + req.path)
							}
							res.status(code)
							res.send()
						})
					}
				} else {
					if(req.query.uploadId !== undefined && req.query.partNumber !== undefined){
						// If it's checking?
						db.all('SELECT hash FROM hashes WHERE id LIKE ?', [req.query.uploadId + "_" + req.query.partNumber], function(err, rows){
							console.log(rows)
							if(rows.hash !== undefined){
								res.status(200)
								res.send()
							} else {
								res.status(500)
								res.send()
							}
						})
					}
				}
			})
		})

		// Handle POST requests
		app.post("/*", function(req, res){
			console.log(req.query)
			if(req.query.uploads !== undefined) {
				// Start a new large upload
				b2_start_large_upload.b2_start_large_upload(auth, req.path, function(code, body){
					if(code == 200) {
						console.log('Starting multi-part upload: ' + body.fileId)
						const ans = '<?xml version="1.0" encoding="UTF-8"?>' +
						'<InitiateMultipartUploadResult xmlns="http://s3.amazonaws.com/doc/2006-03-01/">' +
						'"<UploadId>' + body.fileId + '</UploadId>"' +
						'</InitiateMultipartUploadResult>'
						res.status(200)
						res.send(ans)
					} else {
						res.status(code)
					}
				})
			} else {
				// Finish a large upload
				//console.log(req.query)
				res.status(500)
				res.send(ans)
			}
		})

		// Handle Delete requests
		app.delete("/*", function(req, res){
			b2_delete.b2_delete(auth, req.path, function(code){
				if(code == 200){
					console.log('Deleted file: ' + req.path)
				}
				res.status(code)
				res.send()
			})
		})

		// Return empty Heads
		app.head("/*", function(req, res){
			console.log('Sent empty head')
			res.status(200)
			res.send()
		})

		// Start Listener
		app.listen(9000)
	})
})
