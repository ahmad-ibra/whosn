import { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import { auth } from '../auth/Authorization'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const DetailedEvent = () => {
    const pathname = useLocation().pathname
    const eventID = pathname.substring(pathname.lastIndexOf('/') + 1)

    const [event, setEvent] = useState({})

    useEffect(() => {
        const getEvent = async () => {
            const eventFromServer = await fetchEvent(eventID)
            setEvent(eventFromServer)
        }

        getEvent()
    }, [eventID])

    // Fetch Owned Events
    const fetchEvent = async (id) => {
        const res = await fetch(
            `http://${backendAddress}/api/v1/secured/event/${id}`,
            {
                headers: {
                    'Content-type': 'application/json',
                    Authorization: auth(),
                },
            }
        )

        const data = await res.json()

        return data || {}
    }

    return (
        <div className="container">
            <h2>{event.name}</h2>
            <p>time: {event.time}</p>
            <p>location: {event.location}</p>
            <p>min participants: {event.min_users}</p>
            <p>max participants: {event.max_users}</p>
            <p>price: {event.price}</p>
            <p>link: {event.link}</p>
        </div>
        // TODO: add join/leave event button (depends on if they've joined event or not)
        // TODO: add delete event button (depends on if they're the owner of the event or not)
        // TODO: add share event button (depends on if they're the owner of the event or not)
    )
}

export default DetailedEvent
