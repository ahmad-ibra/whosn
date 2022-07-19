import PropTypes from 'prop-types'
import CustomButton from './CustomButton'

const EventHeader = ({ title, canAddEvent, showAdd, onAdd }) => {
    return (
        <header className="header">
            <h1>{title}</h1>
            {canAddEvent && (
                <CustomButton
                    variant={showAdd ? 'danger' : 'primary'}
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
