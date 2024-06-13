import { useEffect, useState } from 'react'
import styles from './input.module.scss'
import { usePostMessageMutation } from '@/services/api/chat'
import { useAuth0 } from '@auth0/auth0-react'

const Input = () => {
    const [value, setValue] = useState<string>('')
    const [postMessage, {
        isLoading,
    }] = usePostMessageMutation()
    const { user } = useAuth0()

    const handleKeyDown = (e: React.KeyboardEvent) => {
        if (e.key === "Enter") {
            postMessage({ text: value, sentByID: user?.sub || '', sentByName: user?.name || '' })
            setValue('')
        }
    }

    return (
        <div className={styles.Container}>
            <input
                className={styles.Input}
                type='text' placeholder='メッセージを入力'
                value={value}
                onChange={(e) => setValue(e.target.value)}
                onKeyDown={handleKeyDown}
                disabled={isLoading}
            />
        </div>
    )
}

export default Input
