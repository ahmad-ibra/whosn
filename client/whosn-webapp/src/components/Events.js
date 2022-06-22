import PropTypes from 'prop-types'
import SingleEvent from "./SingleEvent"

const Events = ({ events, includeDeleteButton }) => {
    return (
        <>
            {events.map((event) => (
                <SingleEvent key={event.id} event={event} includeDeleteButton={includeDeleteButton} />
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
