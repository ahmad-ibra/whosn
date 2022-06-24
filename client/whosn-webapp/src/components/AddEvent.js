import { useState } from 'react'

const AddEvent = ({ onAdd }) => {
    const [name, setName] = useState('')
    const [time, setTime] = useState('')
    const [location, setLocation] = useState('')
    const [minUsers, setMinUsers] = useState('')
    const [maxUsers, setMaxUsers] = useState('')
    const [price, setPrice] = useState('')

    const onSubmit = (e) => {
        e.preventDefault()

        if (!name || !time || !location || !minUsers || !maxUsers || !price) {
            alert('Please fill out missing event fields')
            return
        }

        // call the function which will write to the backend
        onAdd({ name, time, location })

        setName('')
        setTime('')
        setLocation('')
        setMinUsers(0)
        setMaxUsers(0)
        setPrice(0)
    }

    return (
        <form className='add-form' onSubmit={onSubmit}>
            <div className='form-control'>
                <label>Event</label>
                <input type='text' placeholder='Event Name'
                    value={name} onChange={(e) => setName(e.target.value)} />
            </div>
            <div className='form-control'>
                <label>Date & Time</label>
                <input type='text' placeholder='Monday Jan 1st at 7pm'
                    value={time} onChange={(e) => setTime(e.target.value)} />
            </div>
            <div className='form-control'>
                <label>Location</label>
                <input type='text' placeholder='Location'
                    value={location} onChange={(e) => setLocation(e.target.value)} />
            </div>
            <div className='form-control'>
                <label>Min Attendees</label>
                <input type='number' placeholder='1'
                    value={minUsers} onChange={(e) => setMinUsers(e.target.value)} />
            </div>
            <div className='form-control'>
                <label>Max Attendees</label>
                <input type='number' placeholder='100'
                    value={maxUsers} onChange={(e) => setMaxUsers(e.target.value)} />
            </div>
            <div className='form-control'>
                <label>Price</label>
                <input type='number' step="0.01" placeholder='10.50'
                    value={price} onChange={(e) => setPrice(e.target.value)} />
            </div>

            <input type='submit' value='Save Event' className='btn btn-block' />
        </form>
    )
}
export default AddEvent