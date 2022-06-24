import PropTypes from 'prop-types'
import EventSingle from "./EventSingle"

const Events = ({ events, includeDeleteButton, onDelete }) => {
    return (
        <>
            {events.map((event, index) => (
                <EventSingle key={index} event={event}
                    includeDeleteButton={includeDeleteButton}
                    onDelete={onDelete} />
            ))}
        </>
    )
}

Events.defaultProps = {
    includeDeleteButton: false
}

Events.propTypes = {
    includeDeleteButton: PropTypes.bool,
}

export default Events
