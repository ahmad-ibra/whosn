import PropTypes from 'prop-types'
import { FaTimes } from 'react-icons/fa'

const SingleEvent = ({ event, includeDeleteButton }) => {
    return (
        <div className='event'>
            <h3>
                {event.text}
                {includeDeleteButton &&
                    <FaTimes style={{ color: 'red', cursor: 'pointer' }} />
                }
            </h3>
            <p>{event.day}</p>
        </div>
    )
}

SingleEvent.defaultProps = {
    includeDeleteButton: false
}

SingleEvent.propTypes = {
    includeDeleteButton: PropTypes.bool,
}

export default SingleEvent
