function initMap() {
    var loc = { lat: -6.87948, lng: 107.62258 };
    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 13,
        center: loc,
        gestureHandling: 'cooperative'
    });
    var marker = new google.maps.Marker({
        position: loc,
        map: map
    });
}