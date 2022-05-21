import 'react-native-gesture-handler';

import { Inter_400Regular, Inter_500Medium, useFonts } from '@expo-google-fonts/inter';
import AppLoading from 'expo-app-loading';
import { StatusBar } from 'expo-status-bar';
import { FC } from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { Widget } from './src/components/Widget';
import { theme } from './src/theme';

const App: FC = () => {
    // const [fontsLoaded] = useFonts({ Inter_400Regular, Inter_500Medium });

    // if (!fontsLoaded) return <AppLoading />;

    return (
        <View style={styles.container}>
            <StatusBar style="light" backgroundColor="transparent" translucent />

            <Text style={{ color: '#000' }}>Qualquer coisa</Text>

            <Widget />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: theme.colors.background,
        alignItems: 'center',
        justifyContent: 'center',
    },
});

export default App;
