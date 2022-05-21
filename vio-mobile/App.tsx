import { StatusBar } from 'expo-status-bar';
import { StyleSheet, View } from 'react-native';
import { GetImagesButton } from './src/components/GetImagesButton';

export default function App() {
    return (
        <View style={styles.container}>
            <StatusBar style="dark" translucent />

            <GetImagesButton />
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#333',
        alignItems: 'center',
        justifyContent: 'center',
    },
});
