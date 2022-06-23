import { useState } from 'react'

const AddEvent = ({ onAdd }) => {
    const [text, setText] = useState('')
    const [day, setDay] = useState('')
    const [reminder, setReminder] = useState(false)

    const onSubmit = (e) => {
        e.preventDefault()

        if (!text) {
            alert('Please add an event')
            return
        }

        // call the function which will write to the backend
        onAdd({ text, day, reminder })

        setText('')
        setDay('')
        setReminder(false)
    }

    return (
        <form className='add-form' onSubmit={onSubmit}>
            <div className='form-control'>
                <label>Event</label>
                <input type='text' placeholder='Add Event'
                    value={text} onChange={(e) => setText(e.target.value)} />
            </div>
            <div className='form-control'>
                <label>Date & Time</label>
                <input type='text' placeholder='Add Date & Time'
                    value={day} onChange={(e) => setDay(e.target.value)} />
            </div>
            <div className='form-control form-control-check'>
                <label>Set Reminder</label>
                <input type='checkbox'
                    checked={reminder}
                    value={reminder} onChange={(e) => setReminder(e.currentTarget.checked)} />
            </div>

            <input type='submit' value='Save Event' className='btn btn-block' />
        </form>
    )
}
export default AddEvent