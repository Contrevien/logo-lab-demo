import styles from './message.module.scss'
import cx from 'classnames'

type MessageProp = {
    text: string
    sentBy: string
    createdAt: number
    self?: boolean
}

const Message = ({ text, sentBy, createdAt, self }: MessageProp) => {
    return (
        <div className={cx(styles.Message__Box, {
            [styles.Message__Self]: self,
        })}
        >
            <p className={styles.Message__Name}>{sentBy}</p>
            <p className={styles.Message__Text}>{text}</p>
            <span className={styles.Message__Date}>{createdAt}</span>
        </div>
    )
}

export default Message
