export const RemoveProcotol = (url: string) => {
    return url.replace('http://', '').replace('https://', '')
}
