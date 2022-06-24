import { useState } from 'react'
import PhoneInput from 'react-phone-number-input'

const Register = ({ onAdd }) => {
    const [name, setName] = useState('')
    const [user_name, setUserName] = useState('')
    const [password, setPassWord] = useState('')
    const [email, setEmail] = useState('')
    const [phone_number, setPhoneNumber] = useState()

    const onSubmit = (e) => {
        e.preventDefault()

        if (!name || !user_name || !password || !email || !phone_number) {
            alert('Please fill out missing fields')
            return
        }

        // call the function which will write to the backend
        onAdd({ name, user_name, password, email, phone_number })

        setName('')
        setUserName('')
        setPassWord('')
        setEmail('')
        setPhoneNumber()
    }

    return (
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
                    placeholder="Enter phone number"
                    value={phone_number} onChange={setPhoneNumber} />
            </div>

            <input type='submit' value='Register' className='btn btn-block' />
        </form>
    )
}

export default Register
