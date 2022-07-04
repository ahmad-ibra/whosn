import PropTypes from 'prop-types'
import Button from './Button'

const EventHeader = ({ title, canAddEvent, showAdd, onAdd }) => {
    return (
        <header className="header">
            <h1>{title}</h1>
            {canAddEvent && (
                <Button
                    color={showAdd ? 'red' : 'steelblue'}
                    text={showAdd ? 'Close' : 'Add Event'}
                    onClick={onAdd}
                />
            )}
        </header>
    )
}

EventHeader.defaultProps = {
    title: 'Whosn?',
    canAddEvent: false,
}

EventHeader.propTypes = {
    title: PropTypes.string.isRequired,
    canAddEvent: PropTypes.bool,
}

export default EventHeader
