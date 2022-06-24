import { useState } from 'react'
import { Link } from 'react-router-dom'

const Login = ({ onAdd }) => {
    const [user_name, setUserName] = useState('')
    const [password, setPassWord] = useState('')

    const onSubmit = (e) => {
        e.preventDefault()

        if (!user_name || !password) {
            alert('Please fill out missing fields')
            return
        }

        // call the function which will write to the backend
        onAdd({ user_name, password })

        setUserName('')
        setPassWord('')
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

export default Login
