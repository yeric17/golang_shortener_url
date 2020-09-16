
document.addEventListener("DOMContentLoaded", () => {
	const inputURL = document.getElementById("userURL")
	const mainForm = document.getElementById("mainForm")
	const userResponse = document.getElementById("userResponse")

	mainForm.addEventListener("submit", e => {
		e.preventDefault()
		console.log(inputURL.value)
		var url = "http://localhost:9090/api"
		var data = { url: inputURL.value };

		var misCabeceras = new Headers();
		misCabeceras.set('Content-Type', 'application/json')

		async function getData(){
			let res = await fetch(url, {
				method: 'POST', // or 'PUT'
				body: JSON.stringify(data), // data can be `string` or {object}!
				misCabeceras,
				mode: "cors",
				cache: "default"
			})
	
			let resJson = await res.json()
			console.log(resJson)

			resTemplate = `
			<span class="us_response_item">
				URL: <span>${resJson.long_url}</span>
			</span>
			<span class="us_response_item">
				new URL: <span>http://localhost:9090/${resJson.short_url}</span>
			</span>
			`

			userResponse.innerHTML = resTemplate
		}

		getData()
	})

})
