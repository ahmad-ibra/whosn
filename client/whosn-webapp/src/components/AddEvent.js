import { useState } from 'react'
import { DateTimePicker } from '@progress/kendo-react-dateinputs'
import '@progress/kendo-react-intl'
import '@progress/kendo-react-tooltip'
import '@progress/kendo-react-common'
import '@progress/kendo-react-popup'
import '@progress/kendo-date-math'
import '@progress/kendo-react-dropdowns'

const AddEvent = ({ onAdd }) => {
    const [name, setName] = useState('')
    const [time, setTime] = useState(new Date())
    const [location, setLocation] = useState('')
    const [min_users, setMinUsers] = useState('')
    const [max_users, setMaxUsers] = useState('')
    const [price, setPrice] = useState(0)

    const onSubmit = (e) => {
        e.preventDefault()

        if (!name || !time || !location || !min_users || !max_users) {
            alert('Please fill out missing event fields')
            return
        }

        // call the function which will write to the backend
        onAdd({ name, time, location, min_users, max_users, price })

        setName('')
        setTime('')
        setLocation('')
        setMinUsers(0)
        setMaxUsers(0)
        setPrice(0)
    }

    const defaultDate = new Date()

    return (
        <form className="add-form" onSubmit={onSubmit}>
            <div className="form-control">
                <label>Event</label>
                <input
                    type="text"
                    placeholder="Event Name"
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                />
            </div>
            <div className="form-control">
                <label>Date & Time</label>
                {/* TODO: swap to a free DatTimePicker when i have time */}
                <DateTimePicker
                    format={'dd/MMM/yyyy hh:mm a'}
                    defaultValue={defaultDate}
                    value={time}
                    onChange={(e) => setTime(e.target.value)}
                />
            </div>
            <div className="form-control">
                <label>Location</label>
                <input
                    type="text"
                    placeholder="Location"
                    value={location}
                    onChange={(e) => setLocation(e.target.value)}
                />
            </div>
            <div className="form-control">
                <label>Min Attendees</label>
                <input
                    type="number"
                    placeholder="1"
                    value={min_users}
                    onChange={(e) => setMinUsers(Number(e.target.value))}
                />
            </div>
            <div className="form-control">
                <label>Max Attendees</label>
                <input
                    type="number"
                    placeholder="100"
                    value={max_users}
                    onChange={(e) => setMaxUsers(Number(e.target.value))}
                />
            </div>
            <div className="form-control">
                <label>Price</label>
                <input
                    type="number"
                    step="0.01"
                    placeholder="10.50"
                    value={price}
                    onChange={(e) => setPrice(Number(e.target.value))}
                />
            </div>

            <input type="submit" value="Save Event" className="btn btn-block" />
        </form>
    )
}
export default AddEvent
