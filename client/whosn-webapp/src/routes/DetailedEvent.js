import { useState, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import { auth } from '../auth/Authorization'
import CustomButton from '../components/CustomButton'
import Header from '../components/Header'
import NotFound from './NotFound'
import { RWebShare } from 'react-web-share'
import { BiShareAlt } from 'react-icons/bi'
import ListGroup from 'react-bootstrap/ListGroup'
import Table from 'react-bootstrap/Table'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const toLocalDateTime = (utcDateTime) => {
    let date = new Date(utcDateTime)
    let dateString = date.toString()
    return dateString.substring(0, dateString.indexOf('GMT'))
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

    const togglePayment = async () => {
        const res = await fetch(
            `${backendAddress}/api/v1/secured/event/${curEvent.id}/set_paid`,
            {
                method: 'PUT',
                headers: {
                    'Content-type': 'application/json',
                    Authorization: auth(),
                },
                body: JSON.stringify({ has_paid: !hasPaid }),
            }
        )

        return await res.json()
    }

    const found = !(curEvent.error != null)

    const isJoined =
        participants.filter((p) => p.user_id === curUser.id).length > 0

    const isIn =
        participants.filter((p) => p.user_id === curUser.id && p.is_in === true)
            .length > 0

    const hasPaid =
        participants.filter(
            (p) => p.user_id === curUser.id && p.has_paid === true
        ).length > 0

    const remainingSeats = Math.max(0, curEvent.max_users - participants.length)

    return (
        <div>
            <Header></Header>
            <div>
                {!found && <NotFound />}
                {found && (
                    <div>
                        <div className="container">
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
                            <ListGroup variant="flush">
                                <ListGroup.Item>
                                    {toLocalDateTime(curEvent.time)} at{' '}
                                    {curEvent.location}
                                </ListGroup.Item>
                                <ListGroup.Item>
                                    minimum guests: {curEvent.min_users}
                                </ListGroup.Item>
                                <ListGroup.Item>
                                    maximum guests: {curEvent.max_users}
                                </ListGroup.Item>
                                <ListGroup.Item>
                                    price: ${curEvent.price}
                                </ListGroup.Item>
                                <ListGroup.Item>
                                    {remainingSeats > 0 && (
                                        <p>Only {remainingSeats} spots left!</p>
                                    )}
                                    {remainingSeats <= 0 && (
                                        <p>Get on the wait list!</p>
                                    )}
                                </ListGroup.Item>
                            </ListGroup>
                        </div>

                        <div className="container">
                            <h2>WhosN?</h2>
                            <Table striped bordered hover size="sm">
                                <thead>
                                    <tr>
                                        <th>Name</th>
                                        <th>Joined At</th>
                                        <th>In Waitlist</th>
                                        <th>Has Paid</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {participants.map(
                                        ({
                                            event_id,
                                            user_id,
                                            joined_at,
                                            name,
                                            is_in,
                                            has_paid,
                                        }) => (
                                            <tr>
                                                <td>{name}</td>
                                                <td>
                                                    {toLocalDateTime(joined_at)}
                                                </td>
                                                <td>
                                                    {is_in
                                                        ? 'IN!'
                                                        : 'Waitlist :('}
                                                </td>
                                                <td>
                                                    {has_paid
                                                        ? 'Yes!'
                                                        : 'no...'}
                                                </td>
                                            </tr>
                                        )
                                    )}
                                </tbody>
                            </Table>
                            <CustomButton
                                variant={isJoined ? 'danger' : 'primary'}
                                text={isJoined ? 'Leave Event' : 'Join Event'}
                                onClick={() => {
                                    joinOrLeaveEvent(curEvent.id, isJoined)
                                    refreshPage()
                                }}
                            />
                            {isIn && (
                                <CustomButton
                                    variant="danger"
                                    text="Toggle Payment"
                                    onClick={() => {
                                        togglePayment()
                                        refreshPage()
                                    }}
                                />
                            )}
                        </div>
                    </div>
                )}
            </div>
        </div>
    )
}

export default DetailedEvent
