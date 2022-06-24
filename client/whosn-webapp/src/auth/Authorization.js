export function auth() {
    return localStorage.getItem('jwt')
}

export default auth
