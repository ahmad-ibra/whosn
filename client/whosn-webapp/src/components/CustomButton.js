import PropTypes from 'prop-types'
import Button from 'react-bootstrap/Button'

const CustomButton = ({ size, variant, text, onClick }) => {
    return (
        <Button variant={variant} size={size} onClick={onClick}>
            {text}
        </Button>
    )
}

CustomButton.defaultProps = {
    size: 'sm',
    color: 'steelblue',
    variant: 'primary',
}

CustomButton.propTypes = {
    size: PropTypes.string,
    text: PropTypes.string,
    color: PropTypes.string,
    variant: PropTypes.string,
    onClick: PropTypes.func,
}

export default CustomButton
