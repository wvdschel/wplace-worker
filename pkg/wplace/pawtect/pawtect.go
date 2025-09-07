package pawtect

// Load procedure:
// m.set_user_id(userid)
// X = m.get_load_payload()
// GET /load with body:
//    {"pawtectMe":X,"paint-the":"world","but-not":"using-bots","security":"/.well-known/security.txt"}

// Draw procedure:
// m.set_user_id(userid)
// m.request_url(url-utf8, urllen)
// X = m.get_pawtected_endpoint_payload(payload-data-utf8, payload-data-len);
// POST with body:
//   {"colors":[19],"coords":[349,606],"t": cf-token,"fp":???}
// and headers:
// 	 x-pawtect-token: X
// Values for fp:
//    firefox: 8cd6529df58ff24cc4bf4abe31db9ae0
//    chromium: 20c5d37a4996e3a28c486ea6eecef3c3
