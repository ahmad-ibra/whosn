import PropTypes from 'prop-types'
import EventSingle from "./EventSingle"

const Events = ({ events, includeDeleteButton, onDelete }) => {
    return (
        <>
            {events.map((event) => (
                <EventSingle key={event.id} event={event}
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
