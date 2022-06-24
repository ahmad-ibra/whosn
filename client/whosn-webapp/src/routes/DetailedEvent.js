// TODO: implement the detailed event route
import { useLocation } from "react-router-dom"

const DetailedEvent = () => {
    const location = useLocation()
    console.log(location)

    return (
        <div>
            <p>Detailed events page</p>
        </div>
    )
}

export default DetailedEvent
