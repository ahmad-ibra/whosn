import { useState } from 'react'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { Link } from 'react-router-dom'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const loginUser = async (credentials) => {
    const res = await fetch(`http://${backendAddress}/api/v1/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials),
    })

    return await res.json()
}

const setToken = (userToken) => {
    localStorage.setItem('jwt', userToken)
}

const Login = () => {
    const [searchParams] = useSearchParams()
    const [user_name, setUserName] = useState('')
    const [password, setPassWord] = useState('')

    const navigate = useNavigate()

    const onSubmit = (e) => {
        e.preventDefault()

        if (!user_name || !password) {
            alert('Please fill out missing fields')
            return
        }

        // call the function which will auth with backend
        loginUser({ user_name, password }).then((data) => {
            if ('error' in data) {
                alert('Error logging in')
                return
            }

            setToken(data.token)
            const redirectTo = searchParams.get('redirectTo') || '/'
            navigate(redirectTo)
        })
    }

    return (
        <div className="container">
            <form className="add-form" onSubmit={onSubmit}>
                <div className="form-control">
                    <label>Username</label>
                    <input
                        type="text"
                        placeholder="Username"
                        value={user_name}
                        onChange={(e) => setUserName(e.target.value)}
                    />
                </div>
                <div className="form-control">
                    <label>Password</label>
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassWord(e.target.value)}
                    />
                </div>

                <input type="submit" value="Login" className="btn btn-block" />
            </form>
            <footer>
                <Link to="/register">Create new account</Link>
            </footer>
        </div>
    )
}

export default Login
