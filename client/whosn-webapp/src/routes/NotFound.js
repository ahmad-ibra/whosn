import { Link } from 'react-router-dom'

const NotFound = () => {
    return (
        <div className="container">
            <p>404 - Not Found</p>
            <Link to="/">Go Back to Events</Link>
        </div>
    )
}

export default NotFound
