import { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import { auth } from '../auth/Authorization'
import Button from '../components/Button'
import Header from '../components/Header'
import NotFound from './NotFound'
import { RWebShare } from 'react-web-share'
import { BiShareAlt } from 'react-icons/bi'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const toLocalDateTime = (utcDateTime) => {
    let date = new Date(utcDateTime)
    let dateString = date.toString()
    return dateString.substring(0, dateString.indexOf('('))
}

function refreshPage() {
    window.location.reload(false)
}

const DetailedEvent = () => {
    const pathname = useLocation().pathname
    const eventID = pathname.substring(pathname.lastIndexOf('/') + 1)

    const [curEvent, setEvent] = useState({})
    const [curUser, setUser] = useState({})
    const [participants, setParticipants] = useState([{}])

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
            `${backendAddress}/api/v1/secured/event/${id}`,
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
        const res = await fetch(`${backendAddress}/api/v1/secured/user`, {
            headers: {
                'Content-type': 'application/json',
                Authorization: auth(),
            },
        })

        return await res.json()
    }

    const fetchParticipants = async (id) => {
        const res = await fetch(
            `${backendAddress}/api/v1/secured/event/${id}/users`,
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
        const res = await fetch(
            isJoined
                ? `${backendAddress}/api/v1/secured/event/${id}/leave`
                : `${backendAddress}/api/v1/secured/event/${id}/join`,
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
    const isJoined =
        participants.filter((p) => p.user_id === curUser.id).length > 0
    const remainingSeats = Math.max(0, curEvent.max_users - participants.length)

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
                            <RWebShare
                                data={{
                                    url: window.location.href,
                                    text: `WhosN for ${curEvent.name}?`,
                                    title: `WhosN for ${curEvent.name}?`,
                                }}
                            >
                                <BiShareAlt
                                    size={25}
                                    style={{
                                        cursor: 'pointer',
                                        float: 'right',
                                    }}
                                />
                            </RWebShare>
                            <h2>{curEvent.name}</h2>
                            <p>
                                {toLocalDateTime(curEvent.time)} at{' '}
                                {curEvent.location}{' '}
                            </p>
                            <br />
                            <p>minimum guests: {curEvent.min_users}</p>
                            <p>maximum guests: {curEvent.max_users}</p>
                            <p>price: ${curEvent.price}</p>
                            {remainingSeats > 0 && (
                                <p>Only {remainingSeats} spots left!</p>
                            )}
                            {remainingSeats <= 0 && (
                                <p>Get on the wait list!</p>
                            )}
                        </div>
                        <div className="container">
                            <h2>In</h2>
                            <ol>
                                {/* TODO: Create UserList component which lists all users that are in from fetchJoined */}
                                {participants
                                    .filter(
                                        (participants) =>
                                            participants.is_in === true
                                    )
                                    .map(({ name, user_id }) => (
                                        <li key={user_id}>{name}</li>
                                    ))}
                            </ol>
                        </div>
                        <div className="container">
                            <h2>Wait List</h2>
                            <ol>
                                {/* TODO: Create UserList component which lists all users that are on the waitlist from fetchJoined */}
                                {participants
                                    .filter(
                                        (participants) =>
                                            participants.is_in === false
                                    )
                                    .map(({ name, user_id }) => (
                                        <li key={user_id}>{name}</li>
                                    ))}
                            </ol>
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}

export default DetailedEvent
