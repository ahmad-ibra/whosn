import PropTypes from 'prop-types'
import { BiEdit, BiTrash } from 'react-icons/bi'
import { useNavigate } from 'react-router-dom'

const EventSingle = ({ event, includeDeleteButton, onDelete }) => {
    const navigate = useNavigate()
    return (
        <div className="event">
            <h2>
                {event.name}
                {includeDeleteButton && (
                    <BiTrash
                        style={{
                            color: 'red',
                            cursor: 'pointer',
                            float: 'right',
                        }}
                        onClick={() => onDelete(event.id)}
                    />
                )}
                {
                    <BiEdit
                        style={{ float: 'right' }}
                        onClick={() => {
                            navigate(`/event/${event.id}`)
                        }}
                    />
                }
            </h2>
            <p>{event.location}</p>
            <p>{event.time}</p>
        </div>
    )
}

EventSingle.defaultProps = {
    includeDeleteButton: false,
}

EventSingle.propTypes = {
    includeDeleteButton: PropTypes.bool,
}

export default EventSingle
