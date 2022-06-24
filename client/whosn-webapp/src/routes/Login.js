import { useState } from 'react'
import { Link } from 'react-router-dom'
// import PropTypes from 'prop-types'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const loginUser = async (credentials) => {
    const res = await fetch(`http://${backendAddress}/api/v1/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(credentials)
    })

    const data = await res.json()
    return data
}

const Login = () => {
    const [user_name, setUserName] = useState('')
    const [password, setPassWord] = useState('')

    const onSubmit = (e) => {
        e.preventDefault()

        if (!user_name || !password) {
            alert('Please fill out missing fields')
            return
        }

        // call the function which will auth with backend
        const token = loginUser({ user_name, password })
        console.log(`token: ${token}`)
        // setToken(token);
    }

    return (
        <div className="container">
            <form className='add-form' onSubmit={onSubmit}>
                <div className='form-control'>
                    <label>Username</label>
                    <input type='text' placeholder='Username'
                        value={user_name} onChange={(e) => setUserName(e.target.value)} />
                </div>
                <div className='form-control'>
                    <label>Password</label>
                    <input type='password' placeholder='Password'
                        value={password} onChange={(e) => setPassWord(e.target.value)} />
                </div>

                <input type='submit' value='Login' className='btn btn-block' />
            </form>
            <footer>
                <Link to='/register'>Create new account</Link>
            </footer>
        </div>
    )
}

// Login.propTypes = {
//     setToken: PropTypes.func.isRequired
// }

export default Login
