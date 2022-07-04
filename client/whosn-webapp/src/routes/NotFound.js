import { Link } from 'react-router-dom'
import Header from '../components/Header'

const NotFound = () => {
    return (
        <div>
            <Header></Header>
            <div className="container">
                <p>404 - Not Found</p>
                <Link to="/">Go Back to Events</Link>
            </div>
        </div>
    )
}

export default NotFound
