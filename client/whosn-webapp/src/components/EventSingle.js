import PropTypes from 'prop-types'
import { FaTimes } from 'react-icons/fa'

const EventSingle = ({ event, includeDeleteButton, onDelete }) => {
    return (
        <div className='event'>
            <h3>
                {event.name}
                {includeDeleteButton &&
                    <FaTimes style={{ color: 'red', cursor: 'pointer' }}
                        onClick={() => onDelete(event.id)} />
                }
            </h3>
            <p>{event.location}</p>
            <p>{event.time}</p>
        </div>
    )
}

EventSingle.defaultProps = {
    includeDeleteButton: false
}

EventSingle.propTypes = {
    includeDeleteButton: PropTypes.bool,
}

export default EventSingle
