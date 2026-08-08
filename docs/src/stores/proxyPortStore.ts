import { persistentAtom } from '@nanostores/persistent'

export const proxyPortStore = persistentAtom('proxyPort', '80')
