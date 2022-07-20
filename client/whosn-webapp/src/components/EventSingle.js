import PropTypes from 'prop-types'
import { useNavigate } from 'react-router-dom'
import Dropdown from 'react-bootstrap/Dropdown'
import DropdownButton from 'react-bootstrap/DropdownButton'
import ListGroup from 'react-bootstrap/ListGroup'

const toLocalDateTime = (utcDateTime) => {
    var date = new Date(utcDateTime).toString()
    return date.substring(0, date.indexOf('GMT'))
}

const EventSingle = ({ event, includeDeleteButton, onDelete }) => {
    const navigate = useNavigate()
    return (
        <div className="event">
            <h2> {event.name} </h2>
            <DropdownButton
                title=""
                align="end"
                id="dropdown-menu-align-end"
                style={{ float: 'right' }}
            >
                <Dropdown.Item
                    onClick={() => {
                        navigate(`/event/${event.id}`)
                    }}
                >
                    View
                </Dropdown.Item>
                {includeDeleteButton && (
                    <Dropdown.Item onClick={() => onDelete(event.id)}>
                        Delete
                    </Dropdown.Item>
                )}
            </DropdownButton>
            <ListGroup variant="flush">
                <ListGroup.Item>{event.location}</ListGroup.Item>
                <ListGroup.Item>{toLocalDateTime(event.time)}</ListGroup.Item>
            </ListGroup>
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
