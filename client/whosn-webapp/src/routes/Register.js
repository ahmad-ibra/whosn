import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import PhoneInput from 'react-phone-number-input'
import 'react-phone-number-input/style.css'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

// Register User
const registerUser = async (user) => {
    const res = await fetch(`http://${backendAddress}/api/v1/user`, {
        method: 'POST',
        headers: { 'Content-type': 'application/json' },
        body: JSON.stringify(user)
    })

    return await res.json()
}

const Register = () => {
    const [name, setName] = useState('')
    const [user_name, setUserName] = useState('')
    const [password, setPassWord] = useState('')
    const [email, setEmail] = useState('')
    const [phone_number, setPhoneNumber] = useState()

    const navigate = useNavigate();

    const onSubmit = (e) => {
        e.preventDefault()

        if (!name || !user_name || !password || !email || !phone_number) {
            alert('Please fill out missing fields')
            return
        }

        // call the function which will register user on the backend
        registerUser({ name, user_name, password, email, phone_number }).then(data => {
            if ('error' in data) {
                alert('Error registering')
                return
            }

            navigate('/login');
        })
    }

    return (
        <div className="container">
            <form className='add-form' onSubmit={onSubmit}>
                <div className='form-control'>
                    <label>Name</label>
                    <input type='text' placeholder='Name'
                        value={name} onChange={(e) => setName(e.target.value)} />
                </div>
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
                <div className='form-control'>
                    <label>Email</label>
                    <input type='email' placeholder='email@foo.com'
                        value={email} onChange={(e) => setEmail(e.target.value)} />
                </div>
                <div className='form-control'>
                    <label>Phone Number</label>
                    <PhoneInput
                        defaultCountry='CA'
                        placeholder="Enter phone number"
                        value={phone_number} onChange={setPhoneNumber} />
                </div>

                <input type='submit' value='Register' className='btn btn-block' />
            </form>
        </div>
    )
}

export default Register
