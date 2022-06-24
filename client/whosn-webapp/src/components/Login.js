import { useState } from 'react'

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
        <div>
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
                <a href='/register'>Create new account</a>
            </footer>
        </div>
    )
}

export default Login
