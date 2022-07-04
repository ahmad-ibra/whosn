import { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import { auth } from '../auth/Authorization'
import Button from '../components/Button'
import Header from '../components/Header'
import NotFound from './NotFound'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const toLocalDateTime = (utcDateTime) => {
    var date = new Date(utcDateTime)
    return date.toString()
}

function refreshPage() {
    window.location.reload(false)
}

const DetailedEvent = () => {
    const pathname = useLocation().pathname
    const eventID = pathname.substring(pathname.lastIndexOf('/') + 1)

    const [curEvent, setEvent] = useState({})
    const [participants, setParticipants] = useState([{}])
    const [curUser, setUser] = useState({})

    useEffect(() => {
        const getEvent = async () => {
            const eventFromServer = await fetchEvent(eventID)
            setEvent(eventFromServer)
        }

        const getUser = async () => {
            const userFromServer = await fetchUser()
            setUser(userFromServer)
        }

        const getParticipants = async () => {
            const participantsFromServer = await fetchParticipants(eventID)
            setParticipants(participantsFromServer)
        }

        getEvent()
        getUser()
        getParticipants()
    }, [eventID])

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

    const fetchUser = async () => {
        // const res = await fetch(
        //     `http://${backendAddress}/api/v1/secured/user`,
        //     {
        //         headers: {
        //             'Content-type': 'application/json',
        //             Authorization: auth(),
        //         },
        //     }
        // )

        // return await res.json()
        return {}
    }

    const fetchParticipants = async (id) => {
        const res = await fetch(
            // TODO: need to implement this endpoint on the backend. ideally it returns both the joined and waitlist items
            `http://${backendAddress}/api/v1/secured/event/${id}/users`,
            {
                headers: {
                    'Content-type': 'application/json',
                    Authorization: auth(),
                },
            }
        )

        return await res.json()
    }

    const joinOrLeaveEvent = async (id, isJoined) => {
        console.log('clicked botton')
        const res = await fetch(
            // TODO: need to implement this endpoint on the backend. ideally it returns both the joined and waitlist items
            isJoined
                ? `http://${backendAddress}/api/v1/secured/event/${id}/leave`
                : `http://${backendAddress}/api/v1/secured/event/${id}/join`,
            {
                headers: {
                    'Content-type': 'application/json',
                    Authorization: auth(),
                },
            }
        )

        return await res.json()
    }

    const found = !(curEvent.error != null)

    // const isJoined = () => {
    // return participants.filter((p) => p.id === curUser.id).length > 0
    // }
    let isJoined = false

    return (
        <div>
            <Header></Header>
            <div>
                {!found && <NotFound />}
                {found && (
                    <div>
                        <div className="container">
                            <Button
                                color={isJoined ? 'red' : 'steelblue'}
                                text={isJoined ? 'Leave Event' : 'Join Event'}
                                onClick={() => {
                                    joinOrLeaveEvent(curEvent.id, isJoined)
                                    refreshPage()
                                }}
                            />
                            {/* TODO: add delete event button (depends on if they're the owner of the event or not) */}
                            {/* TODO: add share event button (depends on if they're the owner of the event or not) */}
                            <h2>Event Details</h2>
                            <p>name: {curEvent.name}</p>
                            <p>time: {toLocalDateTime(curEvent.time)}</p>
                            <p>location: {curEvent.location}</p>
                            <p>min participants: {curEvent.min_users}</p>
                            <p>max participants: {curEvent.max_users}</p>
                            <p>price: {curEvent.price}</p>
                            {/* TODO: update this to event.max_users - people that are joined */}
                            <p>
                                spots left:{' '}
                                {Math.max(
                                    0,
                                    curEvent.max_users - participants.length
                                )}
                            </p>
                            <p>link: {curEvent.link + curEvent.id}</p>
                        </div>
                        <div className="container">
                            <h2>In</h2>
                            {/* TODO: Create UserList component which lists all users that are in from fetchJoined */}
                            {typeof participants !== 'undefined' &&
                                participants !== 'null' &&
                                participants
                                    .filter(
                                        (participants) =>
                                            participants.is_in === true
                                    )
                                    .map(({ name, joined_at }) => (
                                        <li key={name}>
                                            {name} -{' '}
                                            {toLocalDateTime(joined_at)}
                                        </li>
                                    ))}
                        </div>
                        <div className="container">
                            <h2>Wait List</h2>
                            {/* TODO: Create UserList component which lists all users that are on the waitlist from fetchJoined */}
                            {typeof participants !== 'undefined' &&
                                participants !== 'null' &&
                                participants
                                    .filter(
                                        (participants) =>
                                            participants.is_in === false
                                    )
                                    .map(({ name, joined_at }) => (
                                        <li key={name}>
                                            {name} -{' '}
                                            {toLocalDateTime(joined_at)}
                                        </li>
                                    ))}
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}

export default DetailedEvent
