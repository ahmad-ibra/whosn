import { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import { auth } from '../auth/Authorization'
import NotFound from './NotFound'

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

        return await res.json()
    }

    // Fetch Joined Event
    // const fetchJoined = async (id) => {
    //     const res = await fetch(
    //         // TODO: need to implement this endpoint on the backend. ideally it returns both the joined and waitlist items
    //         `http://${backendAddress}/api/v1/secured/event_users/${id}`,
    //         {
    //             headers: {
    //                 'Content-type': 'application/json',
    //                 Authorization: auth(),
    //             },
    //         }
    //     )

    //     return await res.json()
    // }

    const found = !(event.error != null)

    return (
        <div>
            {!found && <NotFound />}
            {found && (
                <div>
                    <div className="container">
                        {/* TODO: add join/leave event button (depends on if they've joined event or not) */}
                        {/* TODO: add delete event button (depends on if they're the owner of the event or not) */}
                        {/* TODO: add share event button (depends on if they're the owner of the event or not) */}
                        <h2>Event Details</h2>
                        <p>name: {event.name}</p>
                        <p>time: {event.time}</p>
                        <p>location: {event.location}</p>
                        <p>min participants: {event.min_users}</p>
                        <p>max participants: {event.max_users}</p>
                        <p>price: {event.price}</p>
                        {/* TODO: update this to event.max_users - people that are joined */}
                        <p>spots left: {event.max_users}</p>
                        <p>link: {event.link}</p>
                    </div>
                    <div className="container">
                        <h2>In</h2>
                        <p>TODO: get people that joined</p>
                        {/* TODO: Create UserList component which lists all users that are in from fetchJoined */}
                    </div>
                    <div className="container">
                        <h2>Wait List</h2>
                        <p>TODO: get people that are on the waitlist</p>
                        {/* TODO: Create UserList component which lists all users that are on the waitlist from fetchJoined */}
                    </div>
                </div>
            )}
        </div>
    )
}

export default DetailedEvent
