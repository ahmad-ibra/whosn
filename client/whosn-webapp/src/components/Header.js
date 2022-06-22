import PropTypes from 'prop-types'
import Button from './Button'

const Header = ({ title, canAddEvent }) => {
    const onClick = () => {
        console.log('click')
    }
    return (
        <header className='header'>
            <h1>{title}</h1>
            {canAddEvent &&
                <Button text='Add Event' onClick={onClick} />
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

// CSS in JS
// const headingStyle = {
//     color: 'blue', backgroundColor: 'black'
// }

export default Header