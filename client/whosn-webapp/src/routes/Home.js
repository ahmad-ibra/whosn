import { useState, useEffect } from 'react'
import EventHeader from '../components/EventHeader'
import Events from '../components/Events'
import AddEvent from '../components/AddEvent'
import { auth } from '../auth/Authorization'
import Header from '../components/Header'

const backendAddress = process.env.REACT_APP_BACKEND_ADDRESS

const Home = () => {
    const [showAddEvent, setShowAddEvent] = useState(false)
    const [ownedEvents, setOwnedEvents] = useState([])
    const [joinedEvents, setJoinedEvents] = useState([])

    useEffect(() => {
        const getOwnedEvents = async () => {
            const ownedEventsFromServer = await fetchOwnedEvents()
            setOwnedEvents(ownedEventsFromServer)
        }

        const getJoinedEvents = async () => {
            const joinedEventsFromServer = await fetchJoinedEvents()
            setJoinedEvents(joinedEventsFromServer)
        }

        getOwnedEvents()
        getJoinedEvents()
    }, [])

    // Fetch Owned Events
    const fetchOwnedEvents = async () => {
        const res = await fetch(
            `${backendAddress}/api/v1/secured/events/owned`,
            {
                headers: {
                    'Content-type': 'application/json',
                    Authorization: auth(),
                },
            }
        )

        const data = await res.json()
        return data || []
    }

    // Fetch Joined Events
    const fetchJoinedEvents = async () => {
        const res = await fetch(
            `${backendAddress}/api/v1/secured/events/joined`,
            {
                headers: {
                    'Content-type': 'application/json',
                    Authorization: auth(),
                },
            }
        )
        const data = await res.json()
        return data || []
    }

    // Add Event
    const addEvent = async (singleEvent) => {
        const res = await fetch(`${backendAddress}/api/v1/secured/event`, {
            method: 'POST',
            headers: {
                'Content-type': 'application/json',
                Authorization: auth(),
            },
            body: JSON.stringify(singleEvent),
        })

        const data = await res.json()
        setOwnedEvents([...ownedEvents, data])
    }

    // Delete Event
    const deleteEvent = async (id) => {
        await fetch(`${backendAddress}/api/v1/secured/event/${id}`, {
            method: 'DELETE',
            headers: {
                'Content-type': 'application/json',
                Authorization: auth(),
            },
        })

        setOwnedEvents(ownedEvents.filter((event) => event.id !== id))
    }

    return (
        <div>
            <Header></Header>
            <div className="container">
                <EventHeader
                    title="My Events"
                    canAddEvent={true}
                    showAdd={showAddEvent}
                    onAdd={() => setShowAddEvent(!showAddEvent)}
                />
                {showAddEvent && <AddEvent onAdd={addEvent} />}
                {ownedEvents.length > 0 && (
                    <Events
                        events={ownedEvents}
                        includeDeleteButton={true}
                        onDelete={deleteEvent}
                    />
                )}
                <br />
                <br />
                <EventHeader title="Joined Events" />
                {joinedEvents.length > 0 && <Events events={joinedEvents} />}
            </div>
        </div>
    )
}

export default Home
