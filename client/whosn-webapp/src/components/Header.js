import PropTypes from 'prop-types'
import Button from './Button'

const Header = ({ title, canAddEvent, showAdd, onAdd }) => {
    return (
        <header className='header'>
            <h1>{title}</h1>
            {canAddEvent &&
                <Button
                    color={showAdd ? 'red' : 'steelblue'}
                    text={showAdd ? 'Close' : 'Add'} onClick={onAdd}
                />
            }
        </header>
    )
}

Header.defaultProps = {
    title: 'Whosn?',
    canAddEvent: false
}

Header.propTypes = {
    title: PropTypes.string.isRequired,
    canAddEvent: PropTypes.bool,
}

export default Header