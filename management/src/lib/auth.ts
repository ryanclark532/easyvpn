function getToken():{ token: string| undefined, is_admin: boolean| undefined}{
	if(localStorage){
		return{
			token: localStorage.getItem("token") ?? undefined,
			is_admin: localStorage.getItem("is_admin") === 'true'
		}
	}
	return {
		token: undefined,
		is_admin: undefined
	}
}

function setToken(token: string, is_admin: boolean){
	if(localStorage){
		localStorage.setItem("token", token)
		localStorage.setItem("is_admin", `${is_admin}`);
	}
}

export async function handleLogin(e: any): Promise<boolean> {
	const formData = new FormData(e.target);
	const body: { [key: string]: string } = {};

	formData.forEach((value, key) => {
		body[key] = value.toString();
	});

	const response = await fetch('http://localhost:8080/user/sign-in', {
		body: JSON.stringify(body),
		method: 'POST'
	});
	if (!response.ok) {
		return false
	}
	const json = await response.json()

	if( !json.token && !json.is_admin){
		return false
	}

	setToken(json.token, json.is_admin)
	return true

}
export async function isAuthed(): Promise<boolean>{
	const {token} = getToken()
	if(!token){
		return false
	}
	const body ={
		token: token
	}

	const response = await fetch('http://localhost:8080/user/check-token', {
		body: JSON.stringify(body),
		method: 'POST'
	});
	if (!response.ok) {
		return false
	}

	const json = await response.json();
	return json.token_valid;
}

export async function isAuthedAdmin(): Promise<boolean>{
	const {token} = getToken()
	if(!token){
		return false
	}
	const body ={
		token: token
	}

	console.log(body)

	const response = await fetch('http://localhost:8080/user/check-token', {
		body: JSON.stringify(body),
		method: 'POST'
	});
	if (!response.ok) {
		return false
	}

	const json = await response.json();
	return json.is_admin && json.token_valid;
}